package algorithm

import (
	"fmt"
	"math"
)

func RunBanditAlgorithm(counts, winnings []int64) (index int, err error) {
	if len(counts) != len(winnings) {
		return 0, fmt.Errorf("counts length must be equal winnings length: %w", err)
	}

	sumCounts := sum(counts)
	maxValue := math.Inf(-1)
	for i, count := range counts {
		k := math.Sqrt((2.0 * math.Log(float64(sumCounts))) / float64(count))
		avg := float64(winnings[i]) / float64(count)
		result := avg + k
		if result > maxValue {
			maxValue = result
			index = i
		}
	}

	return index, nil
}

func sum(s []int64) int64 {
	var total int64
	for _, val := range s {
		total += val
	}
	return total
}
