// AUTO-GENERATED: DO NOT EDIT

package services

import (
	"context"
	"google.golang.org/grpc/metadata"

	"github.com/dnk90/adlive/helpers/errors"
	"github.com/dnk90/adlive/pb"
)

func (s *service) Login(context context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	resp, err := s.login(context, req)
	return resp, errors.ToGRPCError(err)
}

func (s *service) GetUserProfile(context context.Context, req *pb.EmptyMessage) (*pb.GetUserProfileResponse, error) {
	md, ok := metadata.FromIncomingContext(context)
	if !ok {
		ll.S.Errorw("not ok")
	} else {
		ll.S.Infow("md", "data", md)
	}
	resp, err := s.getUserProfile(context, req)
	return resp, errors.ToGRPCError(err)
}

func (s *service) UpdateUserProfile(context context.Context, req *pb.UserProfile) (*pb.CommonResponse, error) {
	resp, err := s.updateUserProfile(context, req)
	return resp, errors.ToGRPCError(err)
}
