package storage

import "time"

type Slot struct {
	ID          uint64
	Description string
}

type Banner struct {
	ID          uint64
	Description string
}

type UserGroup struct {
	ID          uint64
	Description string
}

type BannerInfo struct {
	BannerID    uint64
	SlotID      uint64
	UserGroupID uint64
	ShowCount   int64
	ClickCount  int64
}

type BannerStatistic struct {
	BannerID    uint64
	SlotID      uint64
	UserGroupID uint64
	Date        time.Time
}
