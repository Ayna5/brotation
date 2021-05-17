package sql

import (
	"context"
	"fmt"

	"github.com/Ayna5/bannersRotation/internal/app"
)

func (s *Storage) BannersClickStatisticsFilterByDate(ctx context.Context, from int64, to int64) ([]app.BannerStatistic, error) {
	rows, err := s.db.QueryContext(
		ctx,
		`SELECT banner_id, slot_id, user_group_id, extract(epoch from date) date 
			FROM banner_click 
			WHERE extract(epoch from date) >=$1 AND extract(epoch from date) <=$2`,
		from, to,
	)
	if err != nil {
		return nil, fmt.Errorf("can't get shows info: %w", err)
	}
	defer rows.Close()

	var clicks []app.BannerStatistic
	var click app.BannerStatistic
	for rows.Next() {
		err := rows.Scan(&click.BannerID, &click.SlotID, &click.UserGroupID, &click.Date)
		if err != nil {
			return nil, fmt.Errorf("cannot scan: %w", err)
		}
		clicks = append(clicks, click)
	}

	return clicks, nil
}
