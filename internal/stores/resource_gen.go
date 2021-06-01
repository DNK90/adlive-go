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

type ResourceStore struct {
	*gorm.DB
}

// NewResourceStore creates a new instance of ResourceStore, contains whole common functions
// for a service
func NewResourceStore() *ResourceStore {
	return &ResourceStore{config.LoadDB()}
}

func (m *ResourceStore) Save(object *models.Resource) error {
	object.BeforeCreate(m.DB)
	return m.save(object)
}

func (m *ResourceStore) GetByID(id string) (*models.Resource, bool, error) {
	return m.getByID(id)
}

func (m *ResourceStore) UpdateMap(mm map[string]interface{}) error {
	return m.updateMap(mm)
}

func (m *ResourceStore) Add(params ...interface{}) (int, error) {
	if len(params) < 1 {
		return 0, gorm.ErrInvalidData
	}

	name := params[0].(string)
	if m.Exist(name) {
		return 0, errors.DataExisted
	}
	resource := &models.Resource{
		Name:        name,
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
		parent := new(models.Resource)
		existed, err := m.Get(&parent, parentName)
		if err != nil {
			return 0, err
		}
		if !existed {
			return 0, errors.RecordNotFound
		}

		if parent != nil {
			resource.ParentId = parent.Id
		}
	}

	err := m.Save(resource)
	if err != nil {
		return 0, err
	}
	return resource.Id, nil
}

func (m *ResourceStore) Exist(params ...interface{}) bool {
	return m.Model(&models.Resource{}).Select("id").Where("name=?", params[0].(string)).First(&models.Resource{}).Error == nil
}

func (m *ResourceStore) Get(output interface{}, params ...interface{}) (bool, error) {
	err := m.Model(&models.Resource{}).Where("name=?", params[0].(string)).First(output).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (m *ResourceStore) GetParent(id int) ([]int, error) {
	results := make([]int, 0)
	role, exist, err := m.GetByID(fmt.Sprintf("%d", id))
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, gorm.ErrRecordNotFound
	}
	results = append(results, id)
	for role != nil && role.ParentId > 0 {
		role, _, err = m.GetByID(fmt.Sprintf("%d", role.ParentId))
		if err != nil {
			return nil, err
		}
		results = append(results, role.Id)
	}
	return results, nil
}

