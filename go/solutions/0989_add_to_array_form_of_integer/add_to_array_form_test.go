package add_to_array_form_of_integer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddToArrayForm(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name string
		num  []int
		k    int
		want []int
	}{
		{
			name: "LeetCode Example 1",
			num:  []int{1, 2, 0},
			k:    34,
			want: []int{1, 5, 4},
		},
		{
			name: "LeetCode Example 2",
			num:  []int{2, 7, 4},
			k:    181,
			want: []int{4, 5, 5},
		},
		{
			name: "LeetCode Example 3",
			num:  []int{2, 1, 5},
			k:    806,
			want: []int{1, 0, 2, 1},
		},
		{
			name: "k is zero",
			num:  []int{9, 9},
			k:    0,
			want: []int{9, 9},
		},
		{
			name: "single digit with carry",
			num:  []int{9},
			k:    1,
			want: []int{1, 0},
		},
		{
			name: "multiple carries",
			num:  []int{9, 9, 9},
			k:    1,
			want: []int{1, 0, 0, 0},
		},
		{
			name: "k larger than num",
			num:  []int{0},
			k:    23,
			want: []int{2, 3},
		},
	}

	funcsToTest := map[string]func(num []int, k int) []int{
		"Basic":     addToArrayFormBasic,
		"Optimized": addToArrayForm,
	}

	for fnName, fn := range funcsToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					got := fn(tc.num, tc.k)
					assert.Equal(t, tc.want, got, "Input: num=%v, k=%d", tc.num, tc.k)
				})
			}
		})
	}
}
