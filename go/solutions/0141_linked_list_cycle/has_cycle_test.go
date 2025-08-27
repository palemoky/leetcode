package has_cycle

import (
	"leetcode/go/solutions/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

var funcsToTest = map[string]func(*utils.ListNode) bool{
	"HashMap":   hasCycleHashMap,
	"TwoPoints": hasCycleTwoPoints,
}

func TestHasCycle(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name    string
		nums    []int
		cycleAt int // -1 means no cycle, otherwise the index of the cycle entry
		want    bool
	}{
		{"Cycle in middle", []int{3, 2, 0, -4}, 1, true},
		{"Cycle at head (short)", []int{1, 2}, 0, true},
		{"No cycle", []int{1}, -1, false},
		{"Empty", []int{}, -1, false},
		{"No cycle long", []int{1, 2, 3, 4, 5}, -1, false},
		{"Cycle at head", []int{1, 2, 3}, 0, true},
		{"Cycle at tail", []int{1, 2, 3}, 2, true},
	}

	for fnName, fn := range funcsToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					head := utils.NewCycleList(tc.nums, tc.cycleAt)
					got := fn(head)
					assert.Equal(t, tc.want, got, "%s: input=%v, cycleAt=%d", fnName, tc.nums, tc.cycleAt)
				})
			}
		})
	}
}
