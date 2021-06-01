// OVERRIDE ME

package services

import (
	"context"
	"fmt"
	"github.com/dnk90/adlive/helpers/errors"
	"github.com/dnk90/adlive/internal/stores"
	token2 "github.com/dnk90/adlive/internal/token"
	"google.golang.org/grpc/codes"

	"github.com/dnk90/adlive/pb"
)

func (s *service) login(context context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	var token = &token2.JwtToken{MacAddress: req.Mac}
	if req.Username != "" && req.Password != "" {
		user, exist, err := stores.NewUserStore().GetByUsernamePassword(req.Username, req.Password)
		if err != nil {
			return &pb.LoginResponse{Message: "user/password is incorrect"}, errors.Trace(err)
		}
		if !exist {
			return &pb.LoginResponse{Message: "user not found"}, errors.ErrorGRPC(codes.NotFound, "")
		}
		token.Id = fmt.Sprintf("%d", user.Id)
	} else if req.RefreshToken != "" {
		token.RefreshToken = req.RefreshToken
		if ok := token.ValidateRefreshToken(); !ok {
			return &pb.LoginResponse{
				Message: "Invalid request token",
			}, errors.ErrorGRPC(codes.InvalidArgument, "")
		}
	}
	if err := token.Create(); err != nil {
		return &pb.LoginResponse{Message: "internal error"}, errors.Trace(err)
	}
	return &pb.LoginResponse{
		AccessToken: token.AccessToken,
		RefreshToken: token.RefreshToken,
		ExpiredIn: token.AtExpire(),
	}, nil
}
func (s *service) getUserProfile(context context.Context, req *pb.EmptyMessage) (*pb.GetUserProfileResponse, error) {
	return &pb.GetUserProfileResponse{}, nil
}
func (s *service) updateUserProfile(context context.Context, req *pb.UserProfile) (*pb.CommonResponse, error) {
	return &pb.CommonResponse{}, nil
}
