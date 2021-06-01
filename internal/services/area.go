// OVERRIDE ME

package services

import (
	"context"
	"github.com/dnk90/adlive/helpers/errors"
	"github.com/dnk90/adlive/internal/helpers"
	"github.com/dnk90/adlive/internal/stores"

	"github.com/dnk90/adlive/pb"
)

func (s *service) listArea(context context.Context, req *pb.ListAreaRequest) (*pb.ListAreaResponse, error) {
	ll.S.Infow("start getting areas list", "locationId", req.LocationId)
	userId := helpers.GetUserIdFromContext(context)
	store := stores.NewAreaStore()
	areas, err := store.GetByLocationId(userId, int(req.LocationId))
	if err != nil {
		ll.S.Errorw("error while getting area by locationId", "err", err, "locationId", req.LocationId)
		return &pb.ListAreaResponse{
			Code: helpers.InternalError,
			Message: helpers.MessageFromCode[helpers.InternalError],
		}, errors.Error(errors.Internal, err.Error())
	}
	response := make([]*pb.Area, 0)
	for _, area := range areas {
		response = append(response, &pb.Area{
			AreaId: int32(area.Id),
			Name: area.Name,
			LocationId: int32(area.LocationId),
		})
	}
	return &pb.ListAreaResponse{
		Data: response,
		Code: helpers.Success,
		Message: helpers.MessageFromCode[helpers.Success],
	}, nil
}
func (s *service) removeArea(context context.Context, req *pb.RemoveAreaRequest) (*pb.CommonResponse, error) {
	ll.S.Infow("start removing area", "areaId", req.AreaId)
	userId := helpers.GetUserIdFromContext(context)
	store := stores.NewAreaStore()
	area, existed, err := store.GetByID(int(req.AreaId))
	if err != nil || !existed {
		ll.S.Errorw("error occurred while getting area by ID or area not found", "err", err, "existed", existed)
		return &pb.CommonResponse{
			Code: helpers.RecordNotFound,
			Message: helpers.MessageFromCode[helpers.RecordNotFound],
		}, nil
	}
	if area.OwnerId != userId {
		return &pb.CommonResponse{
			Code: helpers.NotOwner,
			Message: helpers.MessageFromCode[helpers.NotOwner],
		}, nil
	}
	area.Status = int(pb.Status_deactivated)
	if err := store.UpdateMap(area.Id, map[string]interface{}{"status": area.Status}); err != nil {
		ll.S.Errorw("error while updating area", "err", err)
		return &pb.CommonResponse{
			Code: helpers.UpdateAreaFailed,
			Message: helpers.MessageFromCode[helpers.UpdateAreaFailed],
		}, nil
	}
	return &pb.CommonResponse{
		Code: helpers.Success,
		Message: helpers.MessageFromCode[helpers.Success],
	}, nil
}
