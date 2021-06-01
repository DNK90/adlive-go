//AUTO-GENERATED: DO NOT EDIT

package stores

import (
	"github.com/dnk90/adlive/config"
	"github.com/dnk90/adlive/internal/models"
	"github.com/dnk90/adlive/internal/stores/errors"
	"github.com/dnk90/adlive/pb"
	"gorm.io/gorm"
)

type ScreenStore struct {
	*gorm.DB
}

// NewScreenStore creates a new instance of ScreenStore, contains whole common functions
// for a service
func NewScreenStore() *ScreenStore {
	return &ScreenStore{config.LoadDB()}
}

func (m *ScreenStore) Save(object *models.Screen) error {
	object.BeforeCreate(m.DB)
	return m.save(object)
}

func (m *ScreenStore) GetByID(id int, params *models.Params) (*models.Screen, bool, error) {
	return m.getByID(id, params)
}

func (m *ScreenStore) UpdateMap(id int, mm map[string]interface{}) error {
	return m.updateMap(id, mm)
}

func (m *ScreenStore) GetByMacAddress(macAddress string) (*models.Screen, bool, error) {
	screen := new(models.Screen)
	err := m.Model(&models.Screen{}).Where("mac=?", macAddress).First(screen).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, false, nil
		}
		return nil, false, err
	}
	return screen, true, nil
}

func (m *ScreenStore) AddNewScreen(object *models.Screen) error {
	_, existed, err := m.GetByMacAddress(object.MAC)
	if err != nil {
		ll.S.Errorw("error while getting screen by mac address", "err", err)
		return err
	}
	if existed {
		return errors.DataExisted
	}
	// add screen
	err = m.Save(object)
	if err != nil {
		ll.S.Errorw("error while inserting new screen", "err", err)
		return err
	}
	return nil
}

// TODO: apply paging, sorting later
func (m *ScreenStore) GetScreens(userId, locationId, areaId int) ([]*models.Screen, error) {
	var screens []*models.Screen
	db := m.Model(&models.Screen{}).Where("user_id = ? AND status = ?", userId, int(pb.Status_activated))
	if locationId > 0 {
		db = db.Where("location_id = ?", locationId)
	}
	if areaId > 0 {
		db = db.Where("area_id = ?", areaId)
	}
	err := db.Find(&screens).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return make([]*models.Screen, 0), nil
		}
		ll.S.Errorw("error while getting screens", "err", err)
		return nil, err
	}
	return screens, nil
}
