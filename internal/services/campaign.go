// OVERRIDE ME

package services

import (
	"context"
	"github.com/dnk90/adlive/helpers/errors"
	"github.com/dnk90/adlive/internal/helpers"
	"github.com/dnk90/adlive/internal/models"
	"github.com/dnk90/adlive/internal/stores"
	"github.com/dnk90/adlive/pb"
)

func GetVideoNames(videos []models.CampaignVideo) []string {
	results := make([]string, 0)
	for _, video := range videos {
		results = append(results, video.Video.Name)
	}
	return results
}

func (s *service) listCampaign(context context.Context, req *pb.CommonPagingMessage) (*pb.ListCampaignResponse, error) {
	userId := helpers.GetUserIdFromContext(context)
	campaigns, err := stores.NewCampaignStore().GetByUserId(userId, &models.Params{Preloads: map[string][]interface{}{
		"Locations": make([]interface{}, 0),
		"Screens": make([]interface{}, 0),
		"Videos": make([]interface{}, 0),
	}})
	if err != nil {
		return &pb.ListCampaignResponse{
			Code:    helpers.InternalError,
			Message: helpers.MessageFromCode[helpers.InternalError],
		}, errors.Error(errors.Internal, err.Error())
	}
	response := &pb.ListCampaignResponse{Data: make([]*pb.Campaign, 0), Code: helpers.Success, Message: helpers.MessageFromCode[helpers.Success]}
	for _, campaign := range campaigns {
		response.Data = append(response.Data, &pb.Campaign{
			Id:       int32(campaign.Id),
			Name:     campaign.Name,
			FromDate: campaign.StartDate,
			ToDate:   campaign.EndDate,
			Status:   0,
			AdSet: &pb.AdSetBasic{
				TotalVideo: int32(len(campaign.Videos)),
				Played:     int32(campaign.Played),
				Videos:     GetVideoNames(campaign.Videos),
			},
			TotalScreen: int32(len(campaign.Screens)),
		})
	}
	return response, nil
}

func (s *service) getCampaign(context context.Context, req *pb.CampaignRequest) (*pb.CampaignResponse, error) {
	userId := helpers.GetUserIdFromContext(context)
	campaign, existed, err := stores.NewCampaignStore().GetByID(int(req.CampaignId), &models.Params{
		Preloads: map[string][]interface{}{
			"Locations": {
				"location_id IN (SELECT DISTINCT location_id FROM campaign_location)",
			},
			"Screens.Screen": make([]interface{}, 0),
			"Videos.Video": make([]interface{}, 0),
		},
	})
	if err != nil {
		return &pb.CampaignResponse{
			Code:    helpers.InternalError,
			Message: helpers.MessageFromCode[helpers.InternalError],
		}, errors.Error(errors.Internal, err.Error())
	}
	if !existed {
		return &pb.CampaignResponse{
			Code:    helpers.RecordNotFound,
			Message: helpers.MessageFromCode[helpers.RecordNotFound],
		}, errors.Error(errors.NotFound, "")
	}
	if campaign.OwnerId != userId {
		return &pb.CampaignResponse{
			Code:    helpers.InvalidUserId,
			Message: helpers.MessageFromCode[helpers.InvalidUserId],
		}, errors.Error(errors.InvalidArgument, "")
	}
	return &pb.CampaignResponse{
		Data: &pb.Campaign{
			Id:       int32(campaign.Id),
			Name:     campaign.Name,
			FromDate: campaign.StartDate,
			ToDate:   campaign.EndDate,
			Status:   pb.CampaignStatus(campaign.Status),
			AdSet: &pb.AdSetBasic{
				TotalVideo: int32(len(campaign.Videos)),
				Played:     int32(campaign.Played),
				Videos:     GetVideoNames(campaign.Videos),
			},
			TotalScreen: int32(len(campaign.Screens)),
			TotalLocation: int32(len(campaign.Locations)),
		},
		Code:    helpers.Success,
		Message: helpers.MessageFromCode[helpers.Success],
	}, nil
}

func (s *service) editCampaign(context context.Context, req *pb.CampaignDetail) (*pb.CommonResponse, error) {
	userId := helpers.GetUserIdFromContext(context)
	err := stores.NewCampaignStore().UpdateCampaign(userId, req)
	if err != nil {
		return &pb.CommonResponse{
			Code:                 helpers.InternalError,
			Message:              helpers.MessageFromCode[helpers.InternalError],
		}, errors.Error(errors.Internal, err.Error())
	}
	return &pb.CommonResponse{
		Code: helpers.Success,
		Message: helpers.MessageFromCode[helpers.Success],
	}, nil
}
