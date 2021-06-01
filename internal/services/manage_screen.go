// OVERRIDE ME

package services

import (
	"context"
	"github.com/dnk90/adlive/helpers/errors"
	"github.com/dnk90/adlive/internal/helpers"
	"github.com/dnk90/adlive/internal/models"
	"github.com/dnk90/adlive/internal/stores"
	"github.com/dnk90/adlive/pb"
	"time"
)

func (s *service) addScreen(context context.Context, req *pb.AddScreenRequest) (*pb.AddScreenResponse, error) {
	userId := helpers.GetUserIdFromContext(context)
	store := stores.NewScreenStore()
	if userId == 0 {
		return &pb.AddScreenResponse{
			Code: helpers.InvalidUserId,
			Message: helpers.MessageFromCode[helpers.InvalidUserId],
		}, errors.Error(errors.InvalidArgument, "invalid userId")
	}
	// if MacAddress exists and no ownerId selected then return ok
	screen, existed, err := store.GetByMacAddress(req.MacAddress)
	if err != nil {
		ll.S.Errorw("error while getting screen by mac address")
		return &pb.AddScreenResponse{
			Code: helpers.InvalidMACAddress,
			Message: helpers.MessageFromCode[helpers.InvalidMACAddress],
		}, errors.Error(errors.Internal, err.Error())
	}
	if existed {
		// check ownerId
		if screen.OwnerId > 0 && screen.OwnerId != userId {
			return &pb.AddScreenResponse{
				Code: helpers.ScreenAlreadyHasOwnerCode,
				Message: helpers.MessageFromCode[helpers.ScreenAlreadyHasOwnerCode],
			}, nil
		}
	} else {
		// insert screen to DB
		screen = &models.Screen{
			OwnerId:     userId,
			DeviceType:  req.TypeOfDevice,
			MAC:         req.MacAddress,
			OS:          req.Os,
			AppVersion:  req.AppVersion,
			Status:      int(pb.Status_activated),
			CreatedDate: time.Now().Unix(),
			UpdatedDate: time.Now().Unix(),
		}
		if err := store.Save(screen); err != nil {
			ll.S.Errorw("error while saving new screen", "err", err)
			return &pb.AddScreenResponse{
				Code: helpers.InternalError,
				Message: helpers.MessageFromCode[helpers.InternalError],
			}, errors.Error(errors.Internal, err.Error())
		}
	}
	// TODO: Implement notification to device
	return &pb.AddScreenResponse{
		Data:                 &pb.ScreenInfo{
			ScreenId:             int32(screen.Id),
			TypeOfDevice:         screen.DeviceType,
			MacAddress:           screen.MAC,
			Os:                   screen.OS,
			AppVersion:           screen.AppVersion,
			IpAddress:            screen.IPAddress,
		},
		Code: helpers.Success,
		Message: helpers.MessageFromCode[helpers.Success],
	}, nil
}

func (s *service) updateScreen(context context.Context, req *pb.UpdateScreenRequest) (*pb.CommonResponse, error) {
	// get userId
	userId := helpers.GetUserIdFromContext(context)
	if userId == 0 {
		return &pb.CommonResponse{
			Code: helpers.InvalidUserId,
			Message: helpers.MessageFromCode[helpers.InvalidUserId],
		}, errors.Error(errors.InvalidArgument, "invalid userId")
	}
	screenStore := stores.NewScreenStore()
	locationStore := stores.NewLocationStore()
	areaStore := stores.NewAreaStore()

	// get screen from screenId
	screen, existed, err := screenStore.GetByID(int(req.Data.ScreenId), nil)
	if err != nil {
		ll.S.Errorw("error while getting screen by screenId")
		return &pb.CommonResponse{
			Code: helpers.InvalidMACAddress,
			Message: helpers.MessageFromCode[helpers.InvalidMACAddress],
		}, errors.Error(errors.Internal, err.Error())
	}
	if !existed {
		return &pb.CommonResponse{
			Code: helpers.RecordNotFound,
			Message: helpers.MessageFromCode[helpers.RecordNotFound],
		}, errors.Error(errors.NotFound, "")
	}
	if screen.OwnerId > 0 && screen.OwnerId != userId {
		return &pb.CommonResponse{
			Code: helpers.ScreenAlreadyHasOwnerCode,
			Message: helpers.MessageFromCode[helpers.ScreenAlreadyHasOwnerCode],
		}, nil
	}
	// merge data from request to screen
	data := req.Data
	screen.Name = data.ScreenName
	screen.LocationId = int(data.LocationId)
	screen.AreaId = int(data.AreaId)

	// check location info
	_, existed, err = locationStore.GetByID(screen.LocationId)
	if err != nil || !existed {
		return &pb.CommonResponse{
			Code: helpers.InvalidLocation,
			Message: helpers.MessageFromCode[helpers.InvalidLocation],
		}, nil
	}
	// check area info
	_, existed, err = areaStore.GetByID(screen.AreaId)
	if err != nil || !existed {
		return &pb.CommonResponse{
			Code: helpers.InvalidArea,
			Message: helpers.MessageFromCode[helpers.InvalidArea],
		}, nil
	}
	// update screen data
	if err = screenStore.UpdateMap(screen.Id, screen.ToMap()); err != nil {
		return &pb.CommonResponse{
			Code: helpers.UpdateScreenFailed,
			Message: helpers.MessageFromCode[helpers.UpdateScreenFailed],
		}, errors.Error(errors.Internal, err.Error())
	}
	return &pb.CommonResponse{
		Code: helpers.Success,
		Message: helpers.MessageFromCode[helpers.Success],
	}, nil
}

func (s *service) listScreen(context context.Context, req *pb.ListScreenRequest) (*pb.ListScreenResponse, error) {
	// get user id from incomming context
	userId := helpers.GetUserIdFromContext(context)
	if userId == 0 {
		return &pb.ListScreenResponse{
			Code: helpers.InvalidUserId,
			Message: helpers.MessageFromCode[helpers.InvalidUserId],
		}, errors.Error(errors.InvalidArgument, "invalid userId")
	}
	// get screens data
	screens, err := stores.NewScreenStore().GetScreens(userId, int(req.LocationId), int(req.AreaId))
	if err != nil {
		return &pb.ListScreenResponse{
			Code: helpers.InternalError,
			Message: helpers.MessageFromCode[helpers.InternalError],
		}, errors.Error(errors.Internal, err.Error())
	}
	response := &pb.ListScreenResponse{
		Data: make([]*pb.ScreenInfo, 0),
		Code: helpers.Success,
		Message: helpers.MessageFromCode[helpers.Success],
	}
	locationNames := make(map[int]*models.Location)
	areaNames := make(map[int]string)
	for _, screen := range screens {
		locationName := locationNames[screen.LocationId]
		if locationName == nil {
			// get data from database
			location, existed, err :=  stores.NewLocationStore().GetByID(screen.LocationId)
			if err != nil || !existed {
				ll.S.Error("problem while getting location in list screen function", "err", err, "existed", existed, "locationId", screen.LocationId)
				return &pb.ListScreenResponse{
					Code: helpers.InternalError,
					Message: helpers.MessageFromCode[helpers.InternalError],
				}, errors.Error(errors.Internal, "")
			}
			locationNames[location.Id] = location
			locationName = location
		}
		areaName := areaNames[screen.AreaId]
		if areaName == "" {
			// get data from area
			area, existed, err :=  stores.NewAreaStore().GetByID(screen.AreaId)
			if err != nil || !existed {
				ll.S.Error("problem while getting area in list screen function", "err", err, "existed", existed, "areaId", screen.AreaId)
				return &pb.ListScreenResponse{
					Code: helpers.InternalError,
					Message: helpers.MessageFromCode[helpers.InternalError],
				}, errors.Error(errors.Internal, "")
			}
			areaNames[area.Id] = area.Name
			areaName = area.Name
		}

		// TODO: expose another service to store/get availability devices (caching)
		response.Data = append(response.Data, &pb.ScreenInfo{
			ScreenId:             int32(screen.Id),
			ScreenName:           screen.Name,
			LocationId:           int32(screen.LocationId),
			LocationName:         locationName.Name,
			LocationAddress:      locationName.Address,
			AreaId:               int32(screen.AreaId),
			AreaName:             areaName,
			Status:               pb.Status(screen.Status),
			TypeOfDevice:         screen.DeviceType,
			MacAddress:           screen.MAC,
			Os:                   screen.OS,
			AppVersion:           screen.AppVersion,
			IpAddress:            screen.IPAddress,
		})
	}
	return response, nil
}
