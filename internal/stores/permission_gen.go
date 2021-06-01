//AUTO-GENERATED: DO NOT EDIT

package stores

import (
	"github.com/dnk90/adlive/config"
	"github.com/dnk90/adlive/internal/models"
	"github.com/dnk90/adlive/internal/stores/errors"
	"github.com/dnk90/adlive/pb"
	"gorm.io/gorm"
	"time"
)

type PermissionStore struct {
	*gorm.DB
}

// NewPermissionStore creates a new instance of PermissionStore, contains whole common functions
// for a service
func NewPermissionStore() *PermissionStore {
	return &PermissionStore{config.LoadDB()}
}

func (m *PermissionStore) Save(object *models.Permission) error {
	object.BeforeCreate(m.DB)
	return m.save(object)
}

func (m *PermissionStore) GetByID(id string) (*models.Permission, bool, error) {
	return m.getByID(id)
}

func (m *PermissionStore) UpdateMap(mm map[string]interface{}) error {
	return m.updateMap(mm)
}

func (m *PermissionStore) Add(params ...interface{}) (int, error) {
	if len(params) < 3 {
		return 0, gorm.ErrInvalidData
	}

	roleId := params[0].(int)
	resourceId := params[1].(int)
	allowed := params[2].(int)
	if m.Exist(roleId, resourceId) {
		return 0, errors.DataExisted
	}
	permission := &models.Permission{
		Role:        roleId,
		Resource:    resourceId,
		Allowed:     allowed,
		Status:      int(pb.Status_activated),
		CreatedDate: time.Now().Unix(),
		UpdatedDate: time.Now().Unix(),
	}
	err := m.Save(permission)
	if err != nil {
		return 0, err
	}
	return permission.Id, nil
}

func (m *PermissionStore) Exist(params ...interface{}) bool {
	return m.Model(&models.Permission{}).Select("id").Where("role=? AND resource=?", params[0].(int), params[1].(int)).First(&models.Permission{}).Error == nil
}

func (m *PermissionStore) Get(output interface{}, params ...interface{}) (bool, error) {
	roleId, resourceId, allowed := params[0].(int), params[1].(int), params[2].(int)
	err := m.Model(&models.Permission{}).Where("role=? AND resource=? AND allowed=?", roleId, resourceId, allowed).First(output).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
