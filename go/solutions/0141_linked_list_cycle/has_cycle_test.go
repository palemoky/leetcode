package has_cycle

import (
	"leetcode/go/solutions/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

var algorithms = []struct {
	name string
	fn   func(*utils.ListNode) bool
}{
	{"HashMap", hasCycleHashMap},
	{"TwoPoints", hasCycleTwoPoints},
}

func TestHasCycle(t *testing.T) {
	testCases := []struct {
		name    string
		nums    []int
		cycleAt int // -1 means no cycle, otherwise the index of the cycle entry
		want    bool
	}{
		{"LeetCode Example 1", []int{3, 2, 0, -4}, 1, true},
		{"LeetCode Example 2", []int{1, 2}, 0, true},
		{"No cycle", []int{1}, -1, false},
		{"Empty", []int{}, -1, false},
		{"No cycle long", []int{1, 2, 3, 4, 5}, -1, false},
		{"Cycle at head", []int{1, 2, 3}, 0, true},
		{"Cycle at tail", []int{1, 2, 3}, 2, true},
	}

	for _, algo := range algorithms {
		t.Run(algo.name, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					head := utils.NewCycleList(tc.nums, tc.cycleAt)
					got := algo.fn(head)
					assert.Equal(t, tc.want, got, "%s: input=%v, cycleAt=%d", algo.name, tc.nums, tc.cycleAt)
				})
			}
		})
	}
}
