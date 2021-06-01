// AUTO-GENERATED: DO NOT EDIT

package services

import (
	"context"

	"github.com/dnk90/adlive/helpers/errors"
	"github.com/dnk90/adlive/pb"
)

func (s *service) ListArea(context context.Context, req *pb.ListAreaRequest) (*pb.ListAreaResponse, error) {
	resp, err := s.listArea(context, req)
	return resp, errors.ToGRPCError(err)
}

func (s *service) RemoveArea(context context.Context, req *pb.RemoveAreaRequest) (*pb.CommonResponse, error) {
	resp, err := s.removeArea(context, req)
	return resp, errors.ToGRPCError(err)
}
