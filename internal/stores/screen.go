// OVERRIDE ME

package stores

import (
	"github.com/dnk90/adlive/internal/models"
	"gorm.io/gorm"
)

func (m *ScreenStore) save(object *models.Screen) error {
	return m.Create(object).Error
}

func (m *ScreenStore) getByID(id int, params *models.Params) (*models.Screen, bool, error) {
	var object = &models.Screen{}
	db := m.Model(models.Screen{}).Where("id = ?", id)
	if params != nil {
		db = params.Process(db)
	}
	err := db.First(object).Error
	if err == gorm.ErrRecordNotFound {
		return object, false, nil
	}
	return object, true, err
}

func (m *ScreenStore) updateMap(id int, mm map[string]interface{}) error {
	return m.Model(models.Screen{Id: id}).Updates(mm).Error
}
