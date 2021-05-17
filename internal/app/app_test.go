package app

import (
	"context"
	"testing"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

var (
	ctx    = context.Background()
	logger *logrus.Logger
)

func newStorageMock(t *testing.T) *StorageMock {
	storageMock := NewStorageMock(t)

	return storageMock
}

func TestApp_AddBannerToSlot(t *testing.T) {
	tests := []struct {
		name     string
		bannerID uint64
		slotID   uint64
		err      error
	}{
		{
			name:     "test should ok",
			bannerID: 1,
			slotID:   1,
			err:      nil,
		},
		{
			name:     "test should error",
			bannerID: 1,
			slotID:   0,
			err:      errors.New("cannot add banner to slot"), //nolint:errcheck
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storageMock := newStorageMock(t)
			storageMock.AddBannerToSlotMock.Set(func(ctx context.Context, bannerID uint64, slotID uint64) error {
				return tt.err
			})

			app := &App{
				logger:  logger,
				storage: storageMock,
			}

			err := app.storage.AddBannerToSlot(ctx, tt.bannerID, tt.slotID)
			require.Equal(t, tt.err, err)
		})
	}
}

func TestApp_AddClickForBanner(t *testing.T) {
	tests := []struct {
		name        string
		bannerID    uint64
		slotID      uint64
		userGroupID uint64
		err         error
	}{
		{
			name:        "test should ok",
			bannerID:    1,
			slotID:      1,
			userGroupID: 1,
			err:         nil,
		},
		{
			name:        "test should error",
			bannerID:    1,
			slotID:      0,
			userGroupID: 1,
			err:         errors.New("cannot click for banner"), //nolint:errcheck
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storageMock := newStorageMock(t)
			storageMock.AddClickForBannerMock.Set(func(ctx context.Context, bannerID uint64, slotID uint64, userGroupID uint64) (err error) {
				return tt.err
			})

			app := &App{
				logger:  logger,
				storage: storageMock,
			}

			err := app.storage.AddClickForBanner(ctx, tt.bannerID, tt.slotID, tt.userGroupID)
			require.Equal(t, tt.err, err)
		})
	}
}

func TestApp_GetBannerForSlot(t *testing.T) {
	tests := []struct {
		name        string
		slotID      uint64
		userGroupID uint64
		result      uint64
		err         error
	}{
		{
			name:        "test should ok",
			slotID:      1,
			userGroupID: 1,
			result:      1,
			err:         nil,
		},
		{
			name:        "test should error",
			slotID:      0,
			userGroupID: 1,
			result:      0,
			err:         errors.New("cannot get banner for slot"), //nolint:errcheck
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storageMock := newStorageMock(t)
			storageMock.GetBannerForSlotMock.Set(func(ctx context.Context, slotID uint64, userGroupID uint64) (u1 uint64, err error) {
				return tt.result, tt.err
			})

			app := &App{
				logger:  logger,
				storage: storageMock,
			}

			got, err := app.storage.GetBannerForSlot(ctx, tt.slotID, tt.userGroupID)
			require.Equal(t, tt.result, got)
			require.Equal(t, tt.err, err)
		})
	}
}

func TestApp_RemoveBannerFromSlot(t *testing.T) {
	tests := []struct {
		name     string
		bannerID uint64
		slotID   uint64
		err      error
	}{
		{
			name:     "test should ok",
			bannerID: 1,
			slotID:   1,
			err:      nil,
		},
		{
			name:     "test should error",
			bannerID: 1,
			slotID:   0,
			err:      errors.New("cannot remove banner from slot"), //nolint:errcheck
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storageMock := newStorageMock(t)
			storageMock.RemoveBannerFromSlotMock.Set(func(ctx context.Context, bannerID uint64, slotID uint64) (err error) {
				return tt.err
			})

			app := &App{
				logger:  logger,
				storage: storageMock,
			}

			err := app.storage.RemoveBannerFromSlot(ctx, tt.bannerID, tt.slotID)
			require.Equal(t, tt.err, err)
		})
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name    string
		want    *App
		wantErr bool
	}{
		{
			name: "test new app should ok",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var log *logrus.Logger
			storageMock := newStorageMock(t)
			got, err := New(storageMock, log)
			require.NotNil(t, got)
			require.NoError(t, err)
		})
	}
}
