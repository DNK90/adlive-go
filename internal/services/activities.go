// OVERRIDE ME

package services

import (
	"context"

	"github.com/dnk90/adlive/pb"
)

func (s *service) listActivities(context context.Context, req *pb.ListActivitiesRequest) (*pb.ListActivitiesResponse, error) {
	return &pb.ListActivitiesResponse{}, nil
}
