// OVERRIDE ME

package stores

import (
	"github.com/dnk90/adlive/internal/models"
	"gorm.io/gorm"
)

func (m *LocationStore) save(object *models.Location) error {
	return m.Create(object).Error
}

func (m *LocationStore) getByID(id int) (*models.Location, bool, error) {
	var object = &models.Location{}
	err := m.Model(models.Location{}).Where("id = ?", id).First(object).Error
	if err == gorm.ErrRecordNotFound {
		return object, false, nil
	}
	return object, true, err
}

func (m *LocationStore) updateMap(id int, mm map[string]interface{}) error {
	return m.Model(models.Location{Id: id}).Updates(mm).Error
}
