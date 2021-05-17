package logger

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name    string
		level logrus.Level
		path  string
		wantErr bool
	}{
		{
			name: "test for logger should ok",
			level: logrus.InfoLevel,
			path: "./logrus.log",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.level, tt.path)
			require.NotNil(t, got)
			require.NoError(t, err)
		})
	}
}
