package sql

import (
	"context"
	"fmt"
)

func (s *Storage) RemoveBannerFromSlot(ctx context.Context, bannerID, slotID uint64) error {
	_, err := s.db.ExecContext(ctx, "DELETE FROM banner_slot WHERE banner_id=$1 AND slot_id=$2", bannerID, slotID)
	if err != nil {
		return fmt.Errorf("cannot remove from table banner_slot for slot %v and banner %v: %w", slotID, bannerID, err)
	}
	return nil
}
