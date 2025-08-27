package plus_one

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlusOne(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		num  []int
		want []int
	}{
		{
			name: "Simple increment",
			num:  []int{1, 2, 3},
			want: []int{1, 2, 4},
		},
		{
			name: "Increment with carry",
			num:  []int{4, 3, 2, 1},
			want: []int{4, 3, 2, 2},
		},
		{
			name: "Single digit with carry",
			num:  []int{9},
			want: []int{1, 0},
		},
		{
			name: "Big number and multiple carries",
			num:  []int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
			want: []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}

	funcsToTest := map[string]func(digits []int) []int{
		"Naive":     plusOneNaive,
		"Optimized": plusOneOptimized,
	}

	for fnName, fn := range funcsToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					// 由于切片是引用类型，拷贝副本以解决数据竞争
					input := append([]int{}, tc.num...)
					got := fn(input)
					assert.Equal(t, tc.want, got, "Input: num=%v", tc.num)
				})
			}
		})
	}
}
