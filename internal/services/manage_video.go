// OVERRIDE ME

package services

import (
	"context"

	"github.com/dnk90/adlive/pb"
)

func (s *service) listVideo(context context.Context, req *pb.CommonPagingMessage) (*pb.ListVideoResponse, error) {
	return &pb.ListVideoResponse{}, nil
}
func (s *service) getVideo(context context.Context, req *pb.GetVideoRequest) (*pb.GetVideoResponse, error) {
	return &pb.GetVideoResponse{}, nil
}
func (s *service) uploadVideo(context context.Context, req *pb.Video) (*pb.UploadVideoResponse, error) {
	
	return &pb.UploadVideoResponse{}, nil
}
