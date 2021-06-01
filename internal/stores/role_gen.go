//AUTO-GENERATED: DO NOT EDIT

package stores

import (
	"fmt"
	"github.com/dnk90/adlive/config"
	"github.com/dnk90/adlive/internal/models"
	"github.com/dnk90/adlive/internal/stores/errors"
	"github.com/dnk90/adlive/pb"
	"gorm.io/gorm"
	"time"
)

type RoleStore struct {
	*gorm.DB
}

// NewRoleStore creates a new instance of RoleStore, contains whole common functions
// for a service
func NewRoleStore() *RoleStore {
	return &RoleStore{config.LoadDB()}
}

func (m *RoleStore) Save(object *models.Role) error {
	object.BeforeCreate(m.DB)
	return m.save(object)
}

func (m *RoleStore) GetByID(id string) (*models.Role, bool, error) {
	return m.getByID(id)
}

func (m *RoleStore) UpdateMap(mm map[string]interface{}) error {
	return m.updateMap(mm)
}

func (m *RoleStore) Add(params ...interface{}) (int, error) {
	if len(params) < 1 {
		return 0, gorm.ErrInvalidData
	}

	roleName := params[0].(string)
	if m.Exist(roleName) {
		return 0, errors.DataExisted
	}

	role := &models.Role{
		Name:        roleName,
		Status:      int(pb.Status_activated),
		CreatedDate: time.Now().Unix(),
		UpdatedDate: time.Now().Unix(),
	}

	// get parentName
	var parentName string
	if len(params) == 2 {
		parentName = params[1].(string)
	}
	if parentName != "" {
		parent := new(models.Role)
		existed, err := m.Get(&parent, parentName)
		if err != nil {
			return 0, err
		}
		if !existed {
			return 0, errors.RecordNotFound
		}
		if parent != nil {
			role.ParentId = parent.Id
		}
	}

	err := m.Save(role)
	if err != nil {
		return 0, err
	}
	return role.Id, nil
}

func (m *RoleStore) Exist(params ...interface{}) bool {
	return m.Model(&models.Role{}).
		Select("id").Where("name=?", params[0].(string)).
		First(new(models.Role)).Error == nil
}

func (m *RoleStore) Get(output interface{}, params ...interface{}) (bool, error) {
	err := m.Model(&models.Role{}).Where("name=?", params[0].(string)).First(output).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (m *RoleStore) GetParent(roleId int) ([]int, error) {
	results := make([]int, 0)
	role, exist, err := m.GetByID(fmt.Sprintf("%d", roleId))
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, gorm.ErrRecordNotFound
	}
	results = append(results, roleId)
	for role != nil && role.ParentId > 0 {
		role, _, err = m.GetByID(fmt.Sprintf("%d", role.ParentId))
		if err != nil {
			return nil, err
		}
		results = append(results, role.Id)
	}
	return results, nil
}
