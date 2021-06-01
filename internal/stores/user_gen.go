//AUTO-GENERATED: DO NOT EDIT

package stores

import (
	"github.com/dnk90/adlive/config"
	"github.com/dnk90/adlive/internal/models"
	"gorm.io/gorm"
)

type UserStore struct {
	*gorm.DB
}

// NewUserStore creates a new instance of UserStore, contains whole common functions
// for a service
func NewUserStore() *UserStore {
	return &UserStore{config.LoadDB()}
}

func (m *UserStore) Save(object *models.User) error {
	object.BeforeCreate(m.DB)
	return m.save(object)
}

func (m *UserStore) GetByID(id int) (*models.User, bool, error) {
	return m.getByID(id)
}

func (m *UserStore) UpdateMap(mm map[string]interface{}) error {
	return m.updateMap(mm)
}

func (m *UserStore) GetByUsernamePassword(username, password string) (*models.User, bool, error) {
	var object = &models.User{}
	err := m.Model(models.User{}).Where("user_name = ? AND password", username, password).First(object).Error
	if err == gorm.ErrRecordNotFound {
		return object, false, nil
	}
	return object, true, err
}
