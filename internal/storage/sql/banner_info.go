package sql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Ayna5/bannersRotation/internal/storage"
	"github.com/pkg/errors"
)

func (s *Storage) bannerInfo(ctx context.Context, slotID, userGroupID uint64) ([]storage.BannerInfo, error) {
	rows, err := s.db.QueryContext(
		ctx,
		`select bsh.banner_id, bsh.slot_id, bsh.user_group_id, count(bsh.*) show_count, count(bcl.*) click_count from banner_showing bsh
		join (select banner_id, slot_id, user_group_id, count(date) from banner_click group by banner_id, slot_id, user_group_id) bcl
		on (bsh.banner_id = bcl.banner_id and bsh.slot_id = bcl.slot_id and bsh.user_group_id = bcl.user_group_id)
		where bsh.slot_id = $1 and bsh.user_group_id = $2 group by bsh.banner_id, bsh.slot_id, bsh.user_group_id`,
		slotID,
		userGroupID,
	)
	if err != nil {
		return nil, errors.Wrap(err, "cannot execute query")
	}
	defer rows.Close()

	var information []storage.BannerInfo
	var info storage.BannerInfo
	for rows.Next() {
		var clickCount sql.NullInt64
		err := rows.Scan(&info.BannerID, &info.SlotID, &info.UserGroupID, &info.ShowCount, &clickCount)
		if err != nil {
			return nil, fmt.Errorf("cannot scan: %w", err)
		}
		info.ClickCount = clickCount.Int64
		information = append(information, info)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("scan error %w", err)
	}
	if len(information) == 0 {
		return nil, errors.New("banner not found") //nolint:errcheck,govet
	}

	return information, nil
}
