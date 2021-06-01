// AUTO-GENERATED: DO NOT EDIT

package services

import (
	"context"

	"github.com/dnk90/adlive/helpers/errors"
	"github.com/dnk90/adlive/pb"
)

func (s *service) ListCampaign(context context.Context, req *pb.CommonPagingMessage) (*pb.ListCampaignResponse, error) {
	resp, err := s.listCampaign(context, req)
	return resp, errors.ToGRPCError(err)
}

func (s *service) GetCampaign(context context.Context, req *pb.CampaignRequest) (*pb.CampaignResponse, error) {
	resp, err := s.getCampaign(context, req)
	return resp, errors.ToGRPCError(err)
}

func (s *service) EditCampaign(context context.Context, req *pb.CampaignDetail) (*pb.CommonResponse, error) {
	resp, err := s.editCampaign(context, req)
	return resp, errors.ToGRPCError(err)
}
