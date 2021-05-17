package sql

import (
	"context"
	"fmt"
)

func (s *Storage) AddClickForBanner(ctx context.Context, bannerID, slotID, userGroupID uint64) error {
	_, err := s.db.ExecContext(
		ctx,
		"INSERT INTO banner_click (banner_id, slot_id, user_group_id, date) VALUES ($1, $2, $3, current_timestamp)",
		bannerID, slotID, userGroupID,
	)
	if err != nil {
		return fmt.Errorf("can't add view for banner: %w", err)
	}
	return nil
}
