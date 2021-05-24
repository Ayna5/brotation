package sql

import (
	"context"
	"fmt"

	"github.com/Ayna5/bannersRotation/internal/storage"
)

func (s *Storage) AddBannerToSlot(ctx context.Context, bannerID, slotID uint64) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("cannot starts transaction: %w", err)
	}
	defer tx.Rollback() //nolint:errcheck

	_, err = tx.ExecContext(ctx, "INSERT INTO banner_slot (banner_id, slot_id) VALUES ($1, $2)", bannerID, slotID)
	if err != nil {
		return fmt.Errorf("cannot insert to table banner_slot for slot %v and banner %v: %w", slotID, bannerID, err)
	}

	var group storage.UserGroup
	var groups []storage.UserGroup
	rows, err := s.db.QueryContext(ctx, "SELECT id, description FROM user_group")
	if err != nil {
		return fmt.Errorf("cannot get user groups: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&group.ID, &group.Description); err != nil {
			return fmt.Errorf("cannot execute scan: %w", err)
		}
		groups = append(groups, group)
	}
	if err = rows.Err(); err != nil {
		return fmt.Errorf("scan error %w", err)
	}

	for _, group := range groups {
		_, err = tx.ExecContext(ctx, "INSERT INTO banner_showing (banner_id, slot_id, user_group_id, date) VALUES ($1, $2, $3, current_timestamp)", bannerID, slotID, group.ID)
		if err != nil {
			return fmt.Errorf("cannot insert to table banner_showing for slot %v, banner %v, group %v: %w", slotID, bannerID, group.ID, err)
		}
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("cannot execute commit: %w", err)
	}
	return nil
}
