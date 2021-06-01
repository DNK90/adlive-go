// OVERRIDE ME

package stores

import (
	"github.com/dnk90/adlive/internal/models"
	"gorm.io/gorm"
)

func (m *RoleStore) save(object *models.Role) error {
	return m.Create(object).Error
}

func (m *RoleStore) getByID(id string) (*models.Role, bool, error) {
	var object = &models.Role{}
	err := m.Model(models.Role{}).Where("id = ?", id).First(object).Error
	if err == gorm.ErrRecordNotFound {
		return object, false, nil
	}
	return object, true, err
}

func (m *RoleStore) updateMap(mm map[string]interface{}) error {
	return m.Model(models.Role{}).Updates(mm).Error
}
