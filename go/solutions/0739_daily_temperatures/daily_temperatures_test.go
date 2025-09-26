package daily_temperatures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDailyTemperatures(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name  string
		temps []int
		want  []int
	}{
		{
			name:  "mixed_fluctuation",
			temps: []int{73, 74, 75, 71, 69, 72, 76, 73},
			want:  []int{1, 1, 4, 2, 1, 1, 0, 0},
		},
		{
			name:  "strictly_increasing_short",
			temps: []int{30, 40, 50, 60},
			want:  []int{1, 1, 1, 0},
		},
		{
			name:  "sparse_peaks",
			temps: []int{30, 60, 90},
			want:  []int{1, 1, 0},
		},
		{
			name:  "all_same",
			temps: []int{50, 50, 50, 50},
			want:  []int{0, 0, 0, 0},
		},
		{
			name:  "strictly_decreasing",
			temps: []int{100, 90, 80, 70},
			want:  []int{0, 0, 0, 0},
		},
		{
			name:  "strictly_increasing",
			temps: []int{10, 20, 30, 40},
			want:  []int{1, 1, 1, 0},
		},
		{
			name:  "single",
			temps: []int{42},
			want:  []int{0},
		},
		{
			name:  "empty",
			temps: []int{},
			want:  []int{},
		},
	}

	funcsToTest := map[string]func(temperatures []int) []int{
		"BruteForce": dailyTemperaturesBruteForce,
		"StackLTR":   dailyTemperaturesStackLeftToRight,
		"StackRTL":   dailyTemperaturesStackRightToLeft,
	}

	for fnName, fn := range funcsToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range cases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					got := fn(tc.temps)
					assert.Equal(t, tc.want, got)
				})
			}
		})
	}
}
