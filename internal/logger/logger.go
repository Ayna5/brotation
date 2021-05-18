package logger

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

func New(level logrus.Level, path string) (*logrus.Logger, error) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("cannot open file: %v", err)
	}

	logger := logrus.New()
	logger.SetLevel(level)
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logger.SetOutput(f)
	return logger, nil
}
