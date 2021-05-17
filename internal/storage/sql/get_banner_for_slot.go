package sql

import (
	"context"
	"fmt"
	"github.com/Ayna5/bannersRotation/internal/algorithm"

	"github.com/pkg/errors"
)

func (s *Storage) GetBannerForSlot(ctx context.Context, slotID, userGroupID uint64) (uint64, error) {
	results, err := s.bannerInfo(ctx, slotID, userGroupID)
	if err != nil {
		return 0, errors.Wrapf(err, "cannot get information about slot %v and userGroup %v:", slotID, userGroupID)
	}

	showCount := make([]int64, len(results))
	clickCount := make([]int64, len(results))
	for i, s := range results {
		showCount[i] = s.ShowCount
		clickCount[i] = s.ClickCount
	}

	index, err := algorithm.RunBanditAlgorithm(showCount, clickCount)
	if err != nil {
		return 0, fmt.Errorf("cannot execute RunBanditAlgorithm: %w", err)
	}

	bannerID := results[index].BannerID

	err = s.addViewForBanner(ctx, bannerID, slotID, userGroupID)
	if err != nil {
		return 0, fmt.Errorf("cannot add view for banner: %w", err)
	}

	return bannerID, nil
}
