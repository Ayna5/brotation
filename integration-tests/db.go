package test

import (
	"database/sql"
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

func getBannerSlotFromBannerSlotTable(t *testing.T, db *sql.DB, bannerID, slotID uint64) int {
	t.Helper()

	var count int
	rows, err := db.Query("select * from banner_slot where banner_id = $1 and slot_id = $2", bannerID, slotID)
	require.NoError(t, err)
	defer rows.Close()

	for rows.Next() {
		count++
	}
	return count
}

func getBannerClickFromBannerClickTable(t *testing.T, db *sql.DB, bannerID, slotID, userGroupID uint64) int {
	t.Helper()

	var count int
	rows, err := db.Query("select * from banner_click where banner_id = $1 and slot_id = $2 and user_group_id = $3", bannerID, slotID, userGroupID)
	require.NoError(t, err)
	defer rows.Close()

	for rows.Next() {
		count++
	}
	return count
}

func truncate(t *testing.T, db *sql.DB) {
	t.Helper()

	_, err := db.Exec("truncate table banner_click, banner_slot, banner_showing")
	require.NoError(t, err)
}
