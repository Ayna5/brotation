package algorithm

import "testing"

func TestRunBanditAlgorithm(t *testing.T) {
	tests := []struct {
		name      string
		counts    []int64
		winnings  []int64
		wantIndex int
		wantErr   bool
	}{
		{
			name:      "test bandit should ok",
			counts:    []int64{3, 9, 5},
			winnings:  []int64{1, 2, 1},
			wantIndex: 0,
			wantErr:   false,
		},
		{
			name:      "test bandit should err when len(counts) < len(winnings)",
			counts:    []int64{3, 9},
			winnings:  []int64{1, 2, 1},
			wantIndex: 0,
			wantErr:   true,
		},
		{
			name:      "test bandit should err when len(counts) > len(winnings)",
			counts:    []int64{3, 9, 5},
			winnings:  []int64{1, 2},
			wantIndex: 0,
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			gotIndex, err := RunBanditAlgorithm(tt.counts, tt.winnings)
			if (err != nil) != tt.wantErr {
				t.Errorf("RunBanditAlgorithm() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotIndex != tt.wantIndex {
				t.Errorf("RunBanditAlgorithm() gotIndex = %v, want %v", gotIndex, tt.wantIndex)
			}
		})
	}
}

func Test_sum(t *testing.T) {
	tests := []struct {
		name   string
		counts []int64
		want   int64
	}{
		{
			name:   "test sum should zero",
			counts: []int64{},
			want:   0,
		},
		{
			name:   "test sum should ok",
			counts: []int64{2, 8, 4},
			want:   14,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := sum(tt.counts); got != tt.want {
				t.Errorf("sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
