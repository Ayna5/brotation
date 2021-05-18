package test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func NewDB(dsn string) *sql.DB {
	var err error
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		panic(err)
	}
	return db
}

func getBannerSlotFromBannerSlotTable(t *testing.T, db *sql.DB, bannerID, slotID uint64) (int, error) {
	t.Helper()

	var count int
	rows, err := db.Query("select * from banner_slot where banner_id = $1 and slot_id = $2", bannerID, slotID)
	require.NoError(t, err)
	defer rows.Close()

	for rows.Next() {
		count++
	}
	if err = rows.Err(); err != nil {
		return 0, fmt.Errorf("scan error %w", err)
	}
	return count, nil
}

func getBannerClickFromBannerClickTable(t *testing.T, db *sql.DB, bannerID, slotID, userGroupID uint64) (int, error) {
	t.Helper()

	var count int
	rows, err := db.Query("select * from banner_click where banner_id = $1 and slot_id = $2 and user_group_id = $3", bannerID, slotID, userGroupID)
	require.NoError(t, err)
	defer rows.Close()

	for rows.Next() {
		count++
	}
	if err = rows.Err(); err != nil {
		return 0, fmt.Errorf("scan error %w", err)
	}
	return count, nil
}

func truncate(t *testing.T, db *sql.DB) {
	t.Helper()

	_, err := db.Exec("truncate table banner_click, banner_slot, banner_showing")
	require.NoError(t, err)
}
