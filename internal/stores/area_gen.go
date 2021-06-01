//AUTO-GENERATED: DO NOT EDIT

package stores

import (
	"github.com/dnk90/adlive/config"
	"github.com/dnk90/adlive/internal/models"
	"github.com/dnk90/adlive/pb"
	"gorm.io/gorm"
	"time"
)

type AreaStore struct {
	*gorm.DB
}

// NewAreaStore creates a new instance of AreaStore, contains whole common functions
// for a service
func NewAreaStore() *AreaStore {
	return &AreaStore{config.LoadDB()}
}

func (m *AreaStore) Save(object *models.Area) error {
	object.BeforeCreate(m.DB)
	return m.save(object)
}

func (m *AreaStore) GetByID(id int) (*models.Area, bool, error) {
	return m.getByID(id)
}

func (m *AreaStore) UpdateMap(id int, mm map[string]interface{}) error {
	return m.updateMap(id, mm)
}

func (m *AreaStore) GetByLocationIdAndName(locationId int, name string) (result *models.Area, err error) {
	err = m.Model(&models.Area{}).
		Where("location_id=? AND name=?", locationId, name).First(&result).Error
	if err != nil {
		return nil, err
	}
	return
}

func (m *AreaStore) GetByLocationId(userId, locationId int) ([]*models.Area, error) {
	var areas []*models.Area
	err := m.Model(&models.Area{}).Where("user_id=? AND location_id=? AND status=?", userId, locationId, int(pb.Status_activated)).Find(&areas).Error
	return areas, err
}

func (m *AreaStore) Update(area *models.Area) error {
	// validate if ownerId existed or not
	_, existed, err := NewUserStore().GetByID(area.OwnerId)
	if !existed {
		ll.S.Errorw("ownerId does not exist", "err", err, "ownerId", area.OwnerId)
		return gorm.ErrRecordNotFound
	}
	area.UpdatedDate = time.Now().Unix()
	// start update
	return m.updateMap(area.Id, area.ToMap())
}

func (m *AreaStore) Insert(area *models.Area) error {
	_, existed, err := NewLocationStore().GetByID(area.LocationId)
	if !existed {
		ll.S.Errorw("location does not exist", "locationId", area.LocationId)
		return gorm.ErrRecordNotFound
	}
	// validate if ownerId existed or not
	_, existed, err = NewUserStore().GetByID(area.OwnerId)
	if !existed {
		ll.S.Errorw("ownerId does not exist", "err", err, "ownerId", area.OwnerId)
		return gorm.ErrRecordNotFound
	}
	area.CreatedDate = time.Now().Unix()
	area.UpdatedDate = area.CreatedDate

	return m.Save(area)
}

func (m *AreaStore) Upsert(area *models.Area) error {
	var (
		data *models.Area
		err error
		existed bool
	)
	if area.Id > 0 {
		// update area
		data, existed, err = m.GetByID(area.Id)
		if !existed || err != nil {
			if err != nil {
				return err
			}
			return gorm.ErrRecordNotFound
		}
		// start insert
		if area.Name != "" && area.Name != data.Name {
			data.Name = area.Name
		}
		if area.LocationId > 0 && area.LocationId != data.LocationId {
			data.LocationId = area.LocationId
		}
		if area.OwnerId > 0 && area.OwnerId != data.OwnerId {
			data.OwnerId = area.OwnerId
		}
		return m.Update(data)
	}
	// insert
	return m.Insert(area)
}
