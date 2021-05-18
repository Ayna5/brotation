package test

import (
	"context"
	"flag"
	"log"
	"testing"

	"github.com/Ayna5/bannersRotation/internal/logger"
	"github.com/Ayna5/bannersRotation/internal/storage/sql"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

const (
	bannerID    = 1
	slotID      = 1
	userGroupID = 1
)

var dsn = "host=localhost port=6432 user=test password=test dbname=brotation sslmode=disable"

func TestStorage(t *testing.T) {
	var z logrus.Level
	flag.Parse()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logg, err := logger.New(z, "./logrus.log")
	if err != nil {
		log.Fatal("failed to create logger")
	}
	store := sql.New()
	if err := store.Connect(ctx, dsn); err != nil {
		logg.Fatal("failed connection")
	}

	t.Run("Add banner to slot", func(t *testing.T) {
		db := NewDB(dsn)
		truncate(t, db)

		err := store.AddBannerToSlot(ctx, bannerID, slotID)
		require.NoError(t, err)

		cc, err := getBannerSlotFromBannerSlotTable(t, db, bannerID, slotID)
		require.NoError(t, err)
		require.NotEqual(t, 0, cc)
	})

	t.Run("Add click for banner", func(t *testing.T) {
		db := NewDB(dsn)
		truncate(t, db)

		err := store.AddClickForBanner(ctx, bannerID, slotID, userGroupID)
		require.NoError(t, err)

		cc, err := getBannerClickFromBannerClickTable(t, db, bannerID, slotID, userGroupID)
		require.NoError(t, err)
		require.NotEqual(t, 0, cc)
	})

	t.Run("Remove banner from slot", func(t *testing.T) {
		db := NewDB(dsn)
		truncate(t, db)

		err := store.AddBannerToSlot(ctx, bannerID, slotID)
		require.NoError(t, err)

		err = store.RemoveBannerFromSlot(ctx, bannerID, slotID)
		require.NoError(t, err)

		cc, err := getBannerSlotFromBannerSlotTable(t, db, bannerID, slotID)
		require.NoError(t, err)
		require.Equal(t, 0, cc)
	})

	t.Run("Get banner for slot", func(t *testing.T) {
		db := NewDB(dsn)
		truncate(t, db)

		err := store.AddBannerToSlot(ctx, bannerID, slotID)
		require.NoError(t, err)
		err = store.AddClickForBanner(ctx, bannerID, slotID, userGroupID)
		require.NoError(t, err)

		banner, err := store.GetBannerForSlot(ctx, slotID, userGroupID)
		require.NoError(t, err)

		require.Equal(t, uint64(1), banner)
	})
}
