// OVERRIDE ME

package stores

import (
	"github.com/dnk90/adlive/internal/models"
	"gorm.io/gorm"
)

func (m *UserStore) save(object *models.User) error {
	return m.Create(object).Error
}

func (m *UserStore) getByID(id int) (*models.User, bool, error) {
	var object = &models.User{}
	err := m.Model(models.User{}).Where("id = ?", id).First(object).Error
	if err == gorm.ErrRecordNotFound {
		return object, false, nil
	}
	return object, true, err
}

func (m *UserStore) updateMap(mm map[string]interface{}) error {
	return m.Model(models.User{}).Updates(mm).Error
}
