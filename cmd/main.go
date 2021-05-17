package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/Ayna5/bannersRotation/configs"
	"github.com/Ayna5/bannersRotation/internal/app"
	"github.com/Ayna5/bannersRotation/internal/kafka"
	"github.com/Ayna5/bannersRotation/internal/logger"
	"github.com/Ayna5/bannersRotation/internal/server/grpc"
	"github.com/Ayna5/bannersRotation/internal/storage/sql"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "./configs/config.toml", "Path to configuration file")
}

func main() {
	flag.Parse()

	config, err := configs.NewConfig(configFile)
	if err != nil {
		log.Fatalf("cannot get config: %v", err)
	}

	lvl, err := logrus.ParseLevel(config.Logger.Level)
	if err != nil {
		log.Fatalf("cannot parse level: %v", err)
	}
	logg, err := logger.New(lvl, config.Logger.Path)
	if err != nil {
		log.Fatalf("cannot start logger: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	storage := sql.New()
	if err := storage.Connect(ctx, config.DB.Dsn); err != nil {
		logg.Error("cannot connect db")
	}
	defer storage.Close()

	producer, err := kafka.New(config.Kafka.Broker, config.Kafka.Topic)
	if err != nil {
		logg.Fatal("cannot init producer")
	}

	statistic := app.NewStatistic(logg, producer, storage, time.Duration(config.Kafka.Time)*time.Second)

	application, err := app.New(storage, logg)
	if err != nil {
		logg.Error("cannot init app")
	}

	grpc, err := grpc.NewServer(logg, application, config.Server.Port)
	if err != nil {
		logg.Fatal("cannot init grpc server")
	}

	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

		select {
		case <-signals:
			err := grpc.Stop()
			if err != nil {
				logg.Error("cannot close connection", err)
			}
			err = producer.Close()
			if err != nil {
				logg.Error("cannot close connection", err)
			}
			signal.Stop(signals)
		case <-ctx.Done():
		}
	}()

	logg.Info("bannersRotation is running...")
	if err := grpc.Start(); err != nil {
		logg.Fatal("failed to start grpc server")
	}

	logg.Info("statistics is running...")
	statistic.Run(ctx)
}
