package logger

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	tests := struct {
		name  string
		level logrus.Level
		path  string
	}{
		name:  "test for logger should ok",
		level: logrus.InfoLevel,
		path:  "./logrus.log",
	}
	t.Run(tests.name, func(t *testing.T) {
		got, err := New(tests.level, tests.path)
		require.NotNil(t, got)
		require.NoError(t, err)
	})
}
