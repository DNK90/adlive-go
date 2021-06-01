// OVERRIDE ME

package stores

import (
	"github.com/dnk90/adlive/internal/models"
	"gorm.io/gorm"
)

func (m *PermissionStore) save(object *models.Permission) error {
	return m.Create(object).Error
}

func (m *PermissionStore) getByID(id string) (*models.Permission, bool, error) {
	var object = &models.Permission{}
	err := m.Model(models.Permission{}).Where("id = ?", id).First(object).Error
	if err == gorm.ErrRecordNotFound {
		return object, false, nil
	}
	return object, true, err
}

func (m *PermissionStore) updateMap(mm map[string]interface{}) error {
	return m.Model(models.Permission{}).Updates(mm).Error
}
