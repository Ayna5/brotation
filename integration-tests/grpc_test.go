package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"

	banners_rotation_pb "github.com/Ayna5/bannersRotation/pkg/banners-rotation"
)

const (
	host = "localhost:14545"
)

var (
	ctx    = context.Background()
	client banners_rotation_pb.BannersRotationClient
)

func TestServerGRPC(t *testing.T) {
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	client = banners_rotation_pb.NewBannersRotationClient(conn)
	t.Run("Add banner to slot", func(t *testing.T) {
		request := &banners_rotation_pb.AddBannerToSlotRequest{
			BannerId: 1,
			SlotId:   1,
		}
		response, err := client.AddBannerToSlot(ctx, request)

		require.NoError(t, err)
		assert.NotNil(t, response)
	})
	t.Run("Remove banner from slot", func(t *testing.T) {
		request := &banners_rotation_pb.RemoveBannerFromSlotRequest{
			BannerId: 1,
			SlotId:   1,
		}
		response, err := client.RemoveBannerFromSlot(ctx, request)

		require.NoError(t, err)
		assert.NotNil(t, response)
	})
	t.Run("Add click for banner", func(t *testing.T) {
		request := &banners_rotation_pb.AddClickForBannerRequest{
			BannerId:    1,
			SlotId:      1,
			UserGroupId: 1,
		}
		response, err := client.AddClickForBanner(ctx, request)

		require.NoError(t, err)
		assert.NotNil(t, response)
	})
	t.Run("Get banner for slot", func(t *testing.T) {
		request := &banners_rotation_pb.GetBannerForSlotRequest{
			SlotId:      1,
			UserGroupId: 1,
		}
		response, err := client.GetBannerForSlot(ctx, request)

		require.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, uint64(1), response.BannerId)
	})
}
