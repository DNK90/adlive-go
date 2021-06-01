// AUTO-GENERATED: DO NOT EDIT

package services

import (
	"context"

	"github.com/dnk90/adlive/helpers/errors"
	"github.com/dnk90/adlive/pb"
)

func (s *service) AddIssue(context context.Context, req *pb.ReportIssueRequest) (*pb.CommonResponse, error) {
	resp, err := s.addIssue(context, req)
	return resp, errors.ToGRPCError(err)
}
