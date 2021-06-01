// OVERRIDE ME

package stores

import (
	"github.com/dnk90/adlive/internal/models"
	"gorm.io/gorm"
)

func (m *VideoStore) save(object *models.Video) error {
	return m.Create(object).Error
}

func (m *VideoStore) getByID(id int, params *models.Params) (*models.Video, bool, error) {
	var object = &models.Video{}
	db := m.Model(models.Video{}).Where("id = ?", id)
	if params != nil {
		db = params.Process(db)
	}
	err := db.First(object).Error
	if err == gorm.ErrRecordNotFound {
		return object, false, nil
	}
	return object, true, err
}

func (m *VideoStore) updateMap(mm map[string]interface{}) error {
	return m.Model(models.Video{}).Updates(mm).Error
}
