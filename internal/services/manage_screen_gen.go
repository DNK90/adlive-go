// AUTO-GENERATED: DO NOT EDIT

package services

import (
	"context"

	"github.com/dnk90/adlive/helpers/errors"
	"github.com/dnk90/adlive/pb"
)

func (s *service) AddScreen(context context.Context, req *pb.AddScreenRequest) (*pb.AddScreenResponse, error) {
	resp, err := s.addScreen(context, req)
	return resp, errors.ToGRPCError(err)
}

func (s *service) UpdateScreen(context context.Context, req *pb.UpdateScreenRequest) (*pb.CommonResponse, error) {
	resp, err := s.updateScreen(context, req)
	return resp, errors.ToGRPCError(err)
}

func (s *service) ListScreen(context context.Context, req *pb.ListScreenRequest) (*pb.ListScreenResponse, error) {
	resp, err := s.listScreen(context, req)
	return resp, errors.ToGRPCError(err)
}
