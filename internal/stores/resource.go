// OVERRIDE ME

package stores

import (
	"github.com/dnk90/adlive/internal/models"
	"gorm.io/gorm"
)

func (m *ResourceStore) save(object *models.Resource) error {
	return m.Create(object).Error
}

func (m *ResourceStore) getByID(id string) (*models.Resource, bool, error) {
	var object = &models.Resource{}
	err := m.Model(models.Resource{}).Where("id = ?", id).First(object).Error
	if err == gorm.ErrRecordNotFound {
		return object, false, nil
	}
	return object, true, err
}

func (m *ResourceStore) updateMap(mm map[string]interface{}) error {
	return m.Model(models.Resource{}).Updates(mm).Error
}
