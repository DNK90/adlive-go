// OVERRIDE ME

package services

import (
	"context"

	"github.com/dnk90/adlive/pb"
)

func (s *service) addIssue(context context.Context, req *pb.ReportIssueRequest) (*pb.CommonResponse, error) {
	return &pb.CommonResponse{}, nil
}
