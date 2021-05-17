package app

import (
	"context"
	"encoding/json"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	typeShow  string = "show"
	typeClick string = "click"
)

type Statistic struct {
	log      *logrus.Logger
	storage  Storage
	producer producer
	interval time.Duration
}

type producer interface {
	Send(msg []byte) error
	Close() error
}

// NewStatistic creates new statistic
func NewStatistic(log *logrus.Logger, producer producer, storage Storage, time time.Duration) *Statistic {
	return &Statistic{
		log:      log,
		storage:  storage,
		producer: producer,
		interval: time,
	}
}

func startWorker(ctx context.Context, done chan struct{}, interval time.Duration, fn func()) {
	ticker := time.NewTicker(interval)
	for {
		select {
		case <-ctx.Done():
			close(done)
			return
		case <-ticker.C:
			fn()
		}
	}
}

func (s *Statistic) Run(ctx context.Context) {
	doneCh := make(chan struct{})
	go startWorker(ctx, doneCh, s.interval, func() {
		s.sendStatisticMessage(ctx)
	})
	<-doneCh
}

func (s *Statistic) sendStatisticMessage(ctx context.Context) {
	from := time.Now()
	to := from.Add(s.interval)

	shows, err := s.storage.BannersShowStatisticsFilterByDate(ctx, from.Unix(), to.Unix())
	if err != nil {
		s.log.Errorf("cannot get events shows: %v", err)
		return
	}

	clicks, err := s.storage.BannersClickStatisticsFilterByDate(ctx, from.Unix(), to.Unix())
	if err != nil {
		s.log.Errorf("cannot get events clicks: %v", err)
		return
	}

	for _, show := range shows {
		msg, err := json.Marshal(BannerStatistic{
			Type:        typeShow,
			BannerID:    show.BannerID,
			SlotID:      show.SlotID,
			UserGroupID: show.UserGroupID,
			Date:        show.Date,
		})
		if err != nil {
			s.log.Errorf("cannot marshal message for bannerID %v, slotID %v, UserGroupID %v: %v", show.BannerID, show.SlotID, show.UserGroupID, err)
			continue
		}

		err = s.producer.Send(msg)
		if err != nil {
			s.log.Errorf("cannot send event notification: %v", err)
		}
	}

	for _, click := range clicks {
		msg, err := json.Marshal(BannerStatistic{
			Type:        typeClick,
			BannerID:    click.BannerID,
			SlotID:      click.SlotID,
			UserGroupID: click.UserGroupID,
			Date:        click.Date,
		})
		if err != nil {
			s.log.Errorf("cannot marshal message for bannerID %v, slotID %v, UserGroupID %v: %v", click.BannerID, click.SlotID, click.UserGroupID, err)
			continue
		}

		err = s.producer.Send(msg)
		if err != nil {
			s.log.Errorf("cannot send event notification: %v", err)
		}
	}
}

type BannerStatistic struct {
	Type        string
	BannerID    uint64
	SlotID      uint64
	UserGroupID uint64
	Date        float64
}
