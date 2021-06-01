// AUTO-GENERATED: DO NOT EDIT

package services

import (
	"context"

	"github.com/dnk90/adlive/helpers/errors"
	"github.com/dnk90/adlive/pb"
)

func (s *service) ListActivities(context context.Context, req *pb.ListActivitiesRequest) (*pb.ListActivitiesResponse, error) {
	resp, err := s.listActivities(context, req)
	return resp, errors.ToGRPCError(err)
}
