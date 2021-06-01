// AUTO-GENERATED: DO NOT EDIT

package services

import (
	"context"

	"github.com/dnk90/adlive/helpers/errors"
	"github.com/dnk90/adlive/pb"
)

func (s *service) AddLocation(context context.Context, req *pb.AddLocationRequest) (*pb.AddLocationResponse, error) {
	resp, err := s.addLocation(context, req)
	return resp, errors.ToGRPCError(err)
}

func (s *service) Locations(context context.Context, req *pb.CommonPagingMessage) (*pb.ListLocationResponse, error) {
	resp, err := s.locations(context, req)
	return resp, errors.ToGRPCError(err)
}

func (s *service) RemoveLocation(context context.Context, req *pb.RemoveLocationRequest) (*pb.CommonResponse, error) {
	resp, err := s.removeLocation(context, req)
	return resp, errors.ToGRPCError(err)
}
