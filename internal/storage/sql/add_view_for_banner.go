package sql

import (
	"context"
	"fmt"
)

func (s *Storage) addViewForBanner(ctx context.Context, bannerID, slotID, socialID uint64) error {
	_, err := s.db.ExecContext(
		ctx,
		"INSERT INTO banner_showing (banner_id, slot_id, user_group_id, date) VALUES ($1, $2, $3, current_timestamp)",
		bannerID, slotID, socialID,
	)
	if err != nil {
		return fmt.Errorf("cannot add view for banner: %w", err)
	}
	return nil
}

