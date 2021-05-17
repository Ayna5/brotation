package grpc

import (
	"context"
	"fmt"

	banners_rotation_pb "github.com/Ayna5/bannersRotation/pkg/banners-rotation"
)

func (a *Server) AddBannerToSlot(ctx context.Context, req *banners_rotation_pb.AddBannerToSlotRequest) (*banners_rotation_pb.AddBannerToSlotResponse, error) {
	if err := a.api.AddBannerToSlot(ctx, req.GetBannerId(), req.GetSlotId()); err != nil {
		return nil, fmt.Errorf("cannot add banner %v to slot %v: %w", req.GetBannerId(), req.GetSlotId(), err)
	}
	return &banners_rotation_pb.AddBannerToSlotResponse{}, nil
}

func (a *Server) RemoveBannerFromSlot(ctx context.Context, req *banners_rotation_pb.RemoveBannerFromSlotRequest) (*banners_rotation_pb.RemoveBannerFromSlotResponse, error) {
	if err := a.api.RemoveBannerFromSlot(ctx, req.GetBannerId(), req.GetSlotId()); err != nil {
		return nil, fmt.Errorf("cannot remove banner %v to slot %v: %w", req.GetBannerId(), req.GetSlotId(), err)
	}
	return &banners_rotation_pb.RemoveBannerFromSlotResponse{}, nil
}

func (a *Server) AddClickForBanner(ctx context.Context, req *banners_rotation_pb.AddClickForBannerRequest) (*banners_rotation_pb.AddClickForBannerResponse, error) {
	if err := a.api.AddClickForBanner(ctx, req.GetBannerId(), req.GetSlotId(), req.GetUserGroupId()); err != nil {
		return nil, fmt.Errorf("cannot add click for banner %v, slot %v, userGroupID %v: %w", req.GetBannerId(), req.GetSlotId(), req.GetUserGroupId(), err)
	}
	return &banners_rotation_pb.AddClickForBannerResponse{}, nil
}

func (a *Server) GetBannerForSlot(ctx context.Context, req *banners_rotation_pb.GetBannerForSlotRequest) (*banners_rotation_pb.GetBannerForSlotResponse, error) {
	bannerID, err := a.api.GetBannerForSlot(ctx, req.GetSlotId(), req.GetUserGroupId())
	if err != nil {
		return nil, fmt.Errorf("cannot get banner for slot %v, userGroup %v: %w", req.GetSlotId(), req.GetUserGroupId(), err)
	}
	return &banners_rotation_pb.GetBannerForSlotResponse{BannerId: bannerID}, nil
}
