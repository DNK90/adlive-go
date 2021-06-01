//AUTO-GENERATED: DO NOT EDIT

package stores

import (
	"github.com/dnk90/adlive/config"
	"github.com/dnk90/adlive/internal/models"
	"gorm.io/gorm"
)

type LogStore struct {
	*gorm.DB
}

// NewLogStore creates a new instance of LogStore, contains whole common functions
// for a service
func NewLogStore() *LogStore {
	return &LogStore{config.LoadDB()}
}

func (m *LogStore) Save(object *models.Log) error {
	object.BeforeCreate(m.DB)
	return m.save(object)
}

func (m *LogStore) GetByID(id int) (*models.Log, bool, error) {
	return m.getByID(id)
}

func (m *LogStore) UpdateMap(mm map[string]interface{}) error {
	return m.updateMap(mm)
}
