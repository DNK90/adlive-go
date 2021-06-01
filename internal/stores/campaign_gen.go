//AUTO-GENERATED: DO NOT EDIT

package stores

import (
	"github.com/dnk90/adlive/config"
	"github.com/dnk90/adlive/helpers/errors"
	"github.com/dnk90/adlive/internal/helpers"
	"github.com/dnk90/adlive/internal/models"
	"github.com/dnk90/adlive/pb"
	"gorm.io/gorm"
	"time"
)

type CampaignStore struct {
	*gorm.DB
}

// NewCampaignStore creates a new instance of CampaignStore, contains whole common functions
// for a service
func NewCampaignStore() *CampaignStore {
	return &CampaignStore{config.LoadDB()}
}

func (m *CampaignStore) Save(object *models.Campaign) error {
	object.BeforeCreate(m.DB)
	return m.save(object)
}

func (m *CampaignStore) GetByID(id int, params *models.Params) (*models.Campaign, bool, error) {
	return m.getByID(id, params)
}

func (m *CampaignStore) UpdateMap(id int, mm map[string]interface{}) error {
	return m.updateMap(id, mm)
}

func (m *CampaignStore) GetByUserId(userId int, params *models.Params) (campaigns []*models.Campaign, err error) {
	db := m.Model(&models.Campaign{}).Where("campaign.owner_id=? AND campaign.status=?", userId, int(pb.Status_activated))
	if params != nil {
		db = params.Process(db)
	}
	err = db.Find(&campaigns).Error
	return
}

func (m *CampaignStore) AddCampaignAreas(obj []*models.CampaignLocation) error {
	return m.Model(&models.CampaignLocation{}).Create(obj).Error
}

func (m *CampaignStore) AddCampaignVideos(obj []*models.CampaignVideo) error {
	return m.Model(&models.CampaignVideo{}).Create(obj).Error
}

func (m *CampaignStore) AddCampaignScreens(obj []*models.CampaignScreen) error {
	return m.Model(&models.CampaignScreen{}).Create(obj).Error
}

func (m *CampaignStore) IncreasePlayed(userId, campaignId int) error {
	return m.Model(&models.Campaign{}).Where("owner_id = ? AND id = ?", userId, campaignId).
		UpdateColumn("played", gorm.Expr("played + 1")).Error
}

func (m *CampaignStore) UpdateCampaign(userId int, obj *pb.CampaignDetail) error {
	// get campaign data campaignId
	campaign, existed, err := m.GetByID(int(obj.Id), &models.Params{
		Preloads: map[string][]interface{}{
			"Screens": make([]interface{}, 0),
			"Videos": make([]interface{}, 0),
		},
	})
	if err != nil {
		ll.S.Errorw("error while getting campaign in UpdateCampaign", "err", err, "campaign_id", obj.Id)
		return err
	}
	if !existed {
		return gorm.ErrRecordNotFound
	}
	if campaign.OwnerId != userId {
		return errors.Error(errors.InvalidArgument, helpers.MessageFromCode[helpers.InvalidUserId])
	}
	campaign.Name = obj.Name
	if obj.FromDate > 0 && obj.FromDate != campaign.StartDate {
		campaign.StartDate = obj.FromDate
	}
	if obj.ToDate > 0 && obj.ToDate != campaign.EndDate {
		campaign.EndDate = obj.ToDate
	}
	if obj.Status > 0 && int(obj.Status) != campaign.Status {
		campaign.Status = int(obj.Status)
	}
	campaign.UpdatedDate = time.Now().Unix()
	return m.Transaction(func(tx *gorm.DB) error {
		// update campaign
		var campaignStore = &CampaignStore{tx}
		if err := campaignStore.updateMap(campaign.Id, campaign.ToMap()); err != nil {
			return err
		}
		// update screen campaigns
	 	if err = campaignStore.UpdateScreenCampaigns(campaign, obj); err != nil {
	 		return err
		}
		// update video campaigns
		if err = campaignStore.UpdateVideoCampaigns(campaign, obj); err != nil {
			return err
		}
		return nil
	})
}

func (m *CampaignStore) UpdateVideoCampaigns(campaign *models.Campaign, obj *pb.CampaignDetail) error {
	// update status of campaign_video
	// - get list of video belong to campaign before updated
	videos, err := m.GetCampaignVideosByCampaignId(campaign.Id)
	if err != nil {
		ll.S.Errorw("error while calling GetCampaignVideosByCampaignId", "err", err)
	}
	// - compare screens with incoming screenId to get updateList and insertList
	insertList, updateList, err := m.CompareCampaignVideosWithIncomingData(campaign, videos, obj.AdSet.Data)
	if err != nil {
		ll.S.Errorw("error while calling CompareCampaignVideosWithIncomingData", "err", err)
		return err
	}

	if len(insertList) > 0 {
		err = m.Model(&models.CampaignVideo{}).CreateInBatches(insertList, config.Cfg.MySQL.Batches).Error
		if err != nil {
			ll.S.Errorw("error while insert CampaignVideo", "err", err)
			return err
		}
	}

	if len(updateList) > 0 {
		for _, campaignVideo := range updateList {
			if err = m.UpdateMap(campaignVideo.Id, campaignVideo.ToMap()); err != nil {
				ll.S.Errorw("error while update CampaignVideo", "err", err)
				return err
			}
		}
	}
	return nil
}

func (m *CampaignStore) UpdateScreenCampaigns(campaign *models.Campaign, obj *pb.CampaignDetail) error {
	// update status of campaign_screen
	// - get list of screen belong to campaign before updated
	screens, err := m.GetCampaignScreensByCampaignId(campaign.Id)
	if err != nil {
		ll.S.Errorw("error while calling GetCampaignScreensByCampaignId", "err", err)
	}
	// - compare screens with incoming screenId to get updateList and insertList
	insertList, updateList, err := m.CompareCampaignScreensWithIncomingData(campaign, screens, obj.SelectedScreens)
	if err != nil {
		ll.S.Errorw("error while calling CompareCampaignScreensWithIncomingData", "err", err)
		return err
	}

	if len(insertList) > 0 {
		err = m.Model(&models.CampaignScreen{}).CreateInBatches(insertList, config.Cfg.MySQL.Batches).Error
		if err != nil {
			ll.S.Errorw("error while insert CampaignScreen", "err", err)
			return err
		}
	}

	if len(updateList) > 0 {
		for _, campaignScreen := range updateList {
			if err = m.UpdateMap(campaignScreen.Id, campaignScreen.ToMap()); err != nil {
				ll.S.Errorw("error while update CampaignScreen", "err", err)
				return err
			}
		}
	}
	return nil
}

func (m *CampaignStore) GetCampaignScreensByCampaignId(campaignId int) ([]*models.CampaignScreen, error) {
	var data []*models.CampaignScreen
	err := m.Model(&models.CampaignScreen{}).Where("campaign_id=?", campaignId).Find(&data).Error
	return data, err
}

func (m *CampaignStore) GetCampaignVideosByCampaignId(campaignId int) ([]*models.CampaignVideo, error) {
	var data []*models.CampaignVideo
	err := m.Model(&models.CampaignVideo{}).Where("campaign_id=?", campaignId).Find(&data).Error
	return data, err
}

func (m *CampaignStore) CompareCampaignScreensWithIncomingData(campaign *models.Campaign, screens []*models.CampaignScreen, incomingData []int32) ([]*models.CampaignScreen, []*models.CampaignScreen, error) {
	ll.S.Infow("start CompareCampaignScreensWithIncomingData", "userId", campaign.OwnerId)
	var (
		updateList = make([]*models.CampaignScreen, 0)
		insertList = make([]*models.CampaignScreen, 0)
		appeared = make(map[int32]bool)
		screenStore = NewScreenStore()
	)

	for _, incoming := range incomingData {
		appeared[incoming] = false
		// loop through screens and update updateList if screens are not nil
		if screens != nil {
			for _, old := range screens {
				if int(incoming) == old.Id && old.OwnerId == campaign.OwnerId {
					if old.Status == int(pb.Status_deactivated) {
						old.Status = int(pb.Status_activated)
						updateList = append(updateList, old)
					}
					// add into appeared
					appeared[incoming] = true
					break
				}
			}
		}
		// if incoming id does not appear in data then process to insert list
		if !appeared[incoming] {
			// is screen existed?
			// screen must not have time in campaign's range
			screen, existed, err := screenStore.GetByID(int(incoming), &models.Params{
				Preloads: map[string][]interface{}{
					"Campaigns.Campaign": {
						"campaign.start_date > ? AND campaign.start_date <= ? AND campaign.end_date > ? AND campaign.end_date <= ?",
						campaign.StartDate, campaign.EndDate, campaign.StartDate, campaign.EndDate,
					},
				},
			})
			if err != nil || !existed {
				ll.S.Errorw("screenId does not exist or invalid screenId", "err", err, "existed", existed)
				return nil, nil, helpers.InvalidScreenIdErr
			}
			// is screen belonged to owner?
			if screen.OwnerId != campaign.OwnerId {
				ll.S.Errorw("screen's ownerId and userId are not matched", "ownerId", screen.OwnerId, "userId", campaign.OwnerId)
				return nil, nil, helpers.InvalidUserIdErr
			}
			// is screen available (location and area are set and is not belonged to any campaign at given time)
			if screen.LocationId == 0 || screen.AreaId == 0 || len(screen.Campaigns) > 0 {
				return nil, nil, helpers.ScreenIsNotAvailableErr
			}
			insertList = append(insertList, &models.CampaignScreen{
				OwnerId:     screen.OwnerId,
				CampaignId:  campaign.Id,
				ScreenId:    screen.Id,
				Status:      int(pb.Status_activated),
				CreatedDate: time.Now().Unix(),
			})
		}
	}
	return insertList, updateList, nil
}

func (m *CampaignStore) CompareCampaignVideosWithIncomingData(campaign *models.Campaign, videos []*models.CampaignVideo, incomingData []*pb.AdSet_Video) ([]*models.CampaignVideo, []*models.CampaignVideo, error) {
	ll.S.Infow("start CompareCampaignVideosWithIncomingData", "userId", campaign.OwnerId)
	var (
		updateList = make([]*models.CampaignVideo, 0)
		insertList = make([]*models.CampaignVideo, 0)
		appeared = make(map[int32]bool)
		videoStore = NewVideoStore()
	)

	for _, incoming := range incomingData {
		appeared[incoming.Id] = false
		// loop through videos and update updateList if screens are not nil
		if videos != nil {
			for _, old := range videos {
				if int(incoming.Id) == old.Id && old.OwnerId == campaign.OwnerId {
					if old.Status == int(pb.Status_deactivated) {
						old.Status = int(pb.Status_activated)
						updateList = append(updateList, old)
					}
					// add into appeared
					appeared[incoming.Id] = true
					break
				}
			}
		}
		// if incoming id does not appear in data then process to insert list
		if !appeared[incoming.Id] {
			// is video existed?
			video, existed, err := videoStore.GetByID(int(incoming.Id), nil)
			if err != nil || !existed {
				ll.S.Errorw("videoId does not exist or invalid videoId", "err", err, "existed", existed)
				return nil, nil, helpers.InvalidVideoIdErr
			}
			// is screen belonged to owner?
			if video.OwnerId != campaign.OwnerId {
				ll.S.Errorw("video's ownerId and userId are not matched", "ownerId", video.OwnerId, "userId", campaign.OwnerId)
				return nil, nil, helpers.InvalidUserIdErr
			}
			insertList = append(insertList, &models.CampaignVideo{
				OwnerId:     video.OwnerId,
				CampaignId:  campaign.Id,
				VideoId:     video.Id,
				Status:      int(pb.Status_activated),
				CreatedDate: time.Now().Unix(),
			})
		}
	}
	return insertList, updateList, nil
}

