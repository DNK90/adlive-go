// OVERRIDE ME

package services

import (
	"context"
	"github.com/dnk90/adlive/helpers/errors"
	"github.com/dnk90/adlive/internal/helpers"
	"github.com/dnk90/adlive/internal/stores"
	"github.com/dnk90/adlive/pb"
	"google.golang.org/grpc/codes"
)

func (s *service) addLocation(context context.Context, req *pb.AddLocationRequest) (*pb.AddLocationResponse, error) {
	userId := helpers.GetUserIdFromContext(context)
	id, err := stores.NewLocationStore().Upsert(userId, req)
	if err != nil {
		ll.S.Errorw("error while insert Location to database", "err", err)
		return &pb.AddLocationResponse{
			Code: helpers.InternalError,
			Message: helpers.MessageFromCode[helpers.InternalError],
		}, errors.Error(errors.Internal, err.Error())
	}
	return &pb.AddLocationResponse{
		LocationId: int32(id),
		Code: helpers.Success,
		Message: helpers.MessageFromCode[helpers.Success],
	}, nil
}

func (s *service) locations(context context.Context, req *pb.CommonPagingMessage) (*pb.ListLocationResponse, error) {
	userId := helpers.GetUserIdFromContext(context)
	response, err := stores.NewLocationStore().GetByUserId(userId)
	if err != nil {
		return &pb.ListLocationResponse{
			Code: helpers.InternalError,
			Message: helpers.MessageFromCode[helpers.InternalError],
		}, errors.ErrorGRPC(codes.Internal, err.Error())
	}
	response.Code = helpers.Success
	response.Message = helpers.MessageFromCode[helpers.Success]
	return response, nil
}

func (s *service) removeLocation(context context.Context, req *pb.RemoveLocationRequest) (*pb.CommonResponse, error) {
	userId := helpers.GetUserIdFromContext(context)
	store := stores.NewLocationStore()
	location, existed, err := store.GetByID(int(req.LocationId))
	if err != nil || !existed {
		ll.S.Errorw("error occurred while getting location by ID or location not found",
			"err", err, "existed", existed, "locationId", req.LocationId)
		return &pb.CommonResponse{
			Code: helpers.RecordNotFound,
			Message: helpers.MessageFromCode[helpers.RecordNotFound],
		}, nil
	}
	if location.OwnerId != userId {
		return &pb.CommonResponse{
			Code: helpers.NotOwner,
			Message: helpers.MessageFromCode[helpers.NotOwner],
		}, nil
	}
	if err = store.UpdateMap(location.Id, map[string]interface{}{
		"status": int(pb.Status_deactivated),
	}); err != nil {
		return &pb.CommonResponse{
			Code: helpers.InternalError,
			Message: helpers.MessageFromCode[helpers.InternalError],
		}, errors.ErrorGRPC(codes.Internal, err.Error())
	}
	return &pb.CommonResponse{
		Code: helpers.Success,
		Message: helpers.MessageFromCode[helpers.Success],
	}, nil
}
