//AUTO-GENERATED: DO NOT EDIT

package stores

import (
	"github.com/dnk90/adlive/config"
	"github.com/dnk90/adlive/internal/models"
	"gorm.io/gorm"
)

type VideoStore struct {
	*gorm.DB
}

// NewVideoStore creates a new instance of VideoStore, contains whole common functions
// for a service
func NewVideoStore() *VideoStore {
	return &VideoStore{config.LoadDB()}
}

func (m *VideoStore) Save(object *models.Video) error {
	object.BeforeCreate(m.DB)
	return m.save(object)
}

func (m *VideoStore) GetByID(id int, params *models.Params) (*models.Video, bool, error) {
	return m.getByID(id, params)
}

func (m *VideoStore) UpdateMap(mm map[string]interface{}) error {
	return m.updateMap(mm)
}

func (m *VideoStore) GetVideosFromCampaignId(userId, campaignId int) ([]models.Video, error) {
	var result []models.Video
	err := m.Model(&models.CampaignVideo{}).Where("user_id = ? AND campaign_id = ?", userId, campaignId).
		Preload("Video").Find(&result).Error
	return result, err
}
