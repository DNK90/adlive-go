// AUTO-GENERATED: DO NOT EDIT

package services

import (
	"context"

	"github.com/dnk90/adlive/helpers/errors"
	"github.com/dnk90/adlive/pb"
)

func (s *service) ListVideo(context context.Context, req *pb.CommonPagingMessage) (*pb.ListVideoResponse, error) {
	resp, err := s.listVideo(context, req)
	return resp, errors.ToGRPCError(err)
}

func (s *service) GetVideo(context context.Context, req *pb.GetVideoRequest) (*pb.GetVideoResponse, error) {
	resp, err := s.getVideo(context, req)
	return resp, errors.ToGRPCError(err)
}

func (s *service) UploadVideo(context context.Context, req *pb.Video) (*pb.UploadVideoResponse, error) {
	resp, err := s.uploadVideo(context, req)
	return resp, errors.ToGRPCError(err)
}
