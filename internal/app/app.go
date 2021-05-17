package app

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
)

type App struct {
	logger  *logrus.Logger
	storage Storage
}

func New(storage Storage, log *logrus.Logger) (*App, error) {
	return &App{
		logger:  log,
		storage: storage,
	}, nil
}

func (a App) AddBannerToSlot(ctx context.Context, bannerID, slotID uint64) error {
	if err := a.storage.AddBannerToSlot(ctx, bannerID, slotID); err != nil {
		a.logger.Error("cannot add banner to slot error")
		return fmt.Errorf("cannot add banner to slot: %w", err)
	}
	return nil
}

func (a App) RemoveBannerFromSlot(ctx context.Context, bannerID, slotID uint64) error {
	if err := a.storage.RemoveBannerFromSlot(ctx, bannerID, slotID); err != nil {
		a.logger.Error("cannot remove banner from slot error")
		return fmt.Errorf("cannot remove banner from slot: %w", err)
	}
	return nil
}

func (a App) AddClickForBanner(ctx context.Context, bannerID, slotID, userGroupID uint64) error {
	if err := a.storage.AddClickForBanner(ctx, bannerID, slotID, userGroupID); err != nil {
		a.logger.Error("cannot click for banner error")
		return fmt.Errorf("cannot add click for banner: %w", err)
	}
	return nil
}

func (a App) GetBannerForSlot(ctx context.Context, slotID, userGroupID uint64) (uint64, error) {
	bannerID, err := a.storage.GetBannerForSlot(ctx, slotID, userGroupID)
	if err != nil {
		a.logger.Error("cannot get banner for slot error")
		return 0, fmt.Errorf("cannot get banner for slot: %w", err)
	}
	return bannerID, nil
}

type Storage interface {
	AddBannerToSlot(ctx context.Context, bannerID, slotID uint64) error
	RemoveBannerFromSlot(ctx context.Context, bannerID, slotID uint64) error
	AddClickForBanner(ctx context.Context, bannerID, slotID, userGroupID uint64) error
	GetBannerForSlot(ctx context.Context, slotID, userGroupID uint64) (uint64, error)
	BannersShowStatisticsFilterByDate(ctx context.Context, from int64, to int64) ([]BannerStatistic, error)
	BannersClickStatisticsFilterByDate(ctx context.Context, from int64, to int64) ([]BannerStatistic, error)
}
