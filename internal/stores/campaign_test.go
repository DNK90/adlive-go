package stores

import (
	"fmt"
	"github.com/dnk90/adlive/config"
	"github.com/dnk90/adlive/internal/models"
	"github.com/dnk90/adlive/pb"
	"github.com/stretchr/testify/suite"
	"math/rand"
	"os"
	"testing"
	"time"
)

type CampaignTestSuite struct {
	suite.Suite
}

func (s *CampaignTestSuite) SetupSuite() {}
func (s *CampaignTestSuite) TearDownSuite() {
	config.DropDB()
}

func (s *CampaignTestSuite)Test_1_GetByUserId() {
	os.Remove("data.db")
	// add user
	rand.Seed(time.Now().Unix())
	user := &models.User{
		Name:        "test_user",
		Email:       fmt.Sprintf("testemail%d",rand.Int31n(100)),
		Phone:       fmt.Sprintf("1234%d", rand.Int31n(100)),
		Company:     "asdasd",
		Role:        1,
		Username:    fmt.Sprintf("test1%d", rand.Int31n(100)),
		Password:    "123456",
		Status:      0,
	}
	err := NewUserStore().Save(user)
	s.NoError(err)
	// add campaign
	campaign := &models.Campaign{
		Name:        "testCampaign",
		OwnerId:     user.Id,
		Played:      10,
		Status:      0,
		StartDate:   1234,
		EndDate:     5678,
	}
	campaignStore := NewCampaignStore()
	err = campaignStore.Save(campaign)
	s.NoError(err)

	// add location and area
	locationStore := NewLocationStore()
	locationId, err := locationStore.Upsert(user.Id, &pb.AddLocationRequest{
		Name: "location1",
		Address: "123 abc",
		Areas: []*pb.Area{
			{
				Name: "area1",
			},
		},
	})
	s.NoError(err)

	area, err := NewAreaStore().GetByLocationIdAndName(locationId, "area1")
	s.NoError(err)

	// add campaign location
	campaignLocation := &models.CampaignLocation{
		OwnerId:     user.Id,
		CampaignId:  campaign.Id,
		LocationId:  locationId,
		AreaId:      area.Id,
		Status:      0,
	}
	err = campaignStore.AddCampaignAreas([]*models.CampaignLocation{campaignLocation})
	s.NoError(err)

	// add screen
	screen := &models.Screen{
		Name:        "screen1",
		LocationId:  locationId,
		AreaId:      area.Id,
		OwnerId:     user.Id,
		DeviceToken: fmt.Sprintf("asdvcxcv%d", rand.Int31n(100)),
		DeviceType:  "asdqweqwe",
		MAC:         fmt.Sprintf("asdasd%d",rand.Int31n(100)),
		OS:          "asdasdasdasd",
		AppVersion:  "1.0",
		IPAddress:   "192.168.1.1",
		Status:      0,
	}
	screenStore := NewScreenStore()
	err = screenStore.Save(screen)
	s.NoError(err)

	// add campaign screen
	err = campaignStore.AddCampaignScreens([]*models.CampaignScreen{
		{
			OwnerId:     user.Id,
			CampaignId:  campaign.Id,
			ScreenId:    screen.Id,
			Status:      0,
		},
	})
	s.NoError(err)

	// add video
	video := &models.Video{
		Title:       "test",
		Name:        "name",
		Format:      "",
		URL:         "",
		PictureURL:  "",
		OwnerId:     user.Id,
		Size:        100,
		Status:      0,
	}
	err = NewVideoStore().Save(video)
	s.NoError(err)

	campaignVideos := make([]*models.CampaignVideo, 0)
	campaignVideos = append(campaignVideos, &models.CampaignVideo{
		OwnerId:     user.Id,
		CampaignId:  campaign.Id,
		VideoId:     video.Id,
		Status:      0,
	})
	// add campaign video
	err = campaignStore.AddCampaignVideos(campaignVideos)
	s.NoError(err)

	// get campaign By user Id
	campaigns, err := campaignStore.GetByUserId(user.Id, &models.Params{Preloads: map[string][]interface{}{
		"Locations": make([]interface{}, 0),
		"Screens": make([]interface{}, 0),
		"Videos": make([]interface{}, 0),
	}})
	s.NoError(err)
	println(campaigns)

	campaignById, existed, err := campaignStore.GetByID(campaign.Id, &models.Params{
		Preloads: map[string][]interface{}{
			"Locations": {
				"location_id IN (SELECT DISTINCT location_id FROM campaign_location)",
			},
			"Screens.Screen": make([]interface{}, 0),
			"Videos.Video": make([]interface{}, 0),
		},
	})
	s.NoError(err)
	s.True(true, existed)
	s.Len(campaignById.Locations, 1)

	// get screen and campaigns
	rs, existed, err := screenStore.GetByID(screen.Id, &models.Params{
		Preloads: map[string][]interface{}{
			"Campaigns.Campaign": {
				"campaign.start_date > ? OR campaign.end_date <= ?",
				455, 122,
			},
		},
	})
	s.True(existed)
	s.NoError(err)
	s.Len(rs.Campaigns, 1)
	s.Equal(rs.Campaigns[0].Id, campaign.Id)
}

func TestCampaign(t *testing.T) {
	suite.Run(t, new(CampaignTestSuite))
}
