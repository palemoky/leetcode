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
		{"mixed_fluctuation", []int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
		{"strictly_increasing_short", []int{30, 40, 50, 60}, []int{1, 1, 1, 0}},
		{"sparse_peaks", []int{30, 60, 90}, []int{1, 1, 0}},
		{"all_same", []int{50, 50, 50, 50}, []int{0, 0, 0, 0}},
		{"strictly_decreasing", []int{100, 90, 80, 70}, []int{0, 0, 0, 0}},
		{"strictly_increasing", []int{10, 20, 30, 40}, []int{1, 1, 1, 0}},
		{"single", []int{42}, []int{0}},
		{"empty", []int{}, []int{}},
	}

	funcsToTest := map[string]func(temperatures []int) []int{
		"BruteForce":       dailyTemperaturesBruteForce,
		"StackLeftToRight": dailyTemperaturesStackLeftToRight,
		"StackRightToLeft": dailyTemperaturesStackRightToLeft,
	}

	for fnName, fn := range funcsToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range cases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					got := fn(append([]int{}, tc.temps...))
					assert.Equal(t, tc.want, got)
				})
			}
		})
	}
}
