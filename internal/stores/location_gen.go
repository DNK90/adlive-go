//AUTO-GENERATED: DO NOT EDIT

package stores

import (
	"github.com/dnk90/adlive/config"
	"github.com/dnk90/adlive/internal/models"
	"github.com/dnk90/adlive/pb"
	"gorm.io/gorm"
	"time"
)

type LocationStore struct {
	*gorm.DB
}

// NewLocationStore creates a new instance of LocationStore, contains whole common functions
// for a service
func NewLocationStore() *LocationStore {
	return &LocationStore{config.LoadDB()}
}

func (m *LocationStore) Save(object *models.Location) error {
	object.BeforeCreate(m.DB)
	return m.save(object)
}

func (m *LocationStore) GetByID(id int) (*models.Location, bool, error) {
	return m.getByID(id)
}

func (m *LocationStore) UpdateMap(id int, mm map[string]interface{}) error {
	return m.updateMap(id, mm)
}

func (m *LocationStore) Upsert(userId int, data *pb.AddLocationRequest) (int, error) {
	// validate if ownerId existed or not
	_, existed, err := NewUserStore().GetByID(userId)
	if !existed {
		ll.S.Errorw("ownerId does not exist", "err", err, "ownerId", userId)
		return 0, gorm.ErrRecordNotFound
	}
	var locationId int
	if err := m.Transaction(func(tx *gorm.DB) error {
		location := &models.Location{
			Id:          int(data.LocationId),
			Name:        data.Name,
			Address:     data.Address,
			OwnerId:     userId,
			Status:      int(data.Status),
		}
		var e error
		if location.Id > 0 {
			// update
			e = m.updateMap(location.Id, location.ToMap())
		} else {
			// insert
			e = m.Insert(location)
		}
		if e != nil {
			return e
		}
		for _, areaPb := range data.Areas {
			area := &models.Area{
				Id:          int(areaPb.AreaId),
				Name:        areaPb.Name,
				OwnerId:     userId,
				LocationId:  location.Id,
				Status:      int(areaPb.Status),
			}
			e = NewAreaStore().Upsert(area)
			if e != nil {
				return e
			}
		}
		locationId = location.Id
		return nil
	}); err != nil {
		return 0, err
	}
	return locationId, nil
}

func (m *LocationStore) Insert(data *models.Location) error {
	data.CreatedDate = time.Now().Unix()
	data.UpdatedDate = data.CreatedDate
	return m.Save(data)
}

func (m *LocationStore) GetByUserId(userId int) (*pb.ListLocationResponse, error) {
	var locations []*models.Location
	err := m.Model(&models.Location{}).Where("owner_id=?", userId).Find(&locations).Error
	if err != nil {
		return nil, err
	}
	results := &pb.ListLocationResponse{Data: make([]*pb.Location, 0)}
	for _, location := range locations {
		// get areas
		areas, err := NewAreaStore().GetByLocationId(userId, location.Id)
		if err != nil {
			return nil, err
		}
		pbAreas := make([]*pb.Area, 0)
		for _, area := range areas {
			pbAreas = append(pbAreas, &pb.Area{
				AreaId:               int32(area.Id),
				Name:                 area.Name,
				LocationId:           int32(area.LocationId),
				Status:               pb.Status(area.Status),
			})
		}
		results.Data = append(results.Data, &pb.Location{
			LocationId:           int32(location.Id),
			Name:                 location.Name,
			Address:              location.Address,
			Status:               int32(location.Status),
			Areas:                pbAreas,
		})
	}
	return results, nil
}
