// OVERRIDE ME

package stores

import (
	"github.com/dnk90/adlive/internal/models"
	"github.com/dnk90/adlive/pb"
	"gorm.io/gorm"
)

func (m *CampaignStore) save(object *models.Campaign) error {
	return m.Create(object).Error
}

func (m *CampaignStore) getByID(id int, params *models.Params) (*models.Campaign, bool, error) {
	var object = &models.Campaign{}
	db := m.Model(models.Campaign{}).Where("campaign.id = ? AND campaign.status = ?", id, int(pb.Status_activated))
	if params != nil {
		db = params.Process(db)
	}
	err := db.First(object).Error
	if err == gorm.ErrRecordNotFound {
		return object, false, nil
	}
	return object, true, err
}

func (m *CampaignStore) updateMap(id int, mm map[string]interface{}) error {
	return m.Model(models.Campaign{Id: id}).Updates(mm).Error
}
