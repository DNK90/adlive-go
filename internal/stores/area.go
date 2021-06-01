// OVERRIDE ME

package stores

import (
	"github.com/dnk90/adlive/internal/models"
	"gorm.io/gorm"
)

func (m *AreaStore) save(object *models.Area) error {
	return m.Create(object).Error
}

func (m *AreaStore) getByID(id int) (*models.Area, bool, error) {
	var object = &models.Area{}
	err := m.Model(models.Area{}).Where("id = ?", id).First(object).Error
	if err == gorm.ErrRecordNotFound {
		return object, false, nil
	}
	return object, true, err
}

func (m *AreaStore) updateMap(id int, mm map[string]interface{}) error {
	return m.Model(models.Area{Id: id}).Updates(mm).Error
}
