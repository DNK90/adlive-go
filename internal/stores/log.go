// OVERRIDE ME

package stores

import (
	"github.com/dnk90/adlive/internal/models"
	"gorm.io/gorm"
)

func (m *LogStore) save(object *models.Log) error {
	return m.Create(object).Error
}

func (m *LogStore) getByID(id int) (*models.Log, bool, error) {
	var object = &models.Log{}
	err := m.Model(models.Log{}).Where("id = ?", id).First(object).Error
	if err == gorm.ErrRecordNotFound {
		return object, false, nil
	}
	return object, true, err
}

func (m *LogStore) updateMap(mm map[string]interface{}) error {
	return m.Model(models.Log{}).Updates(mm).Error
}
