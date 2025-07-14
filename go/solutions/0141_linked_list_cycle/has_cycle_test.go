package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var algorithms = []struct {
	name string
	fn   func(*ListNode) bool
}{
	{"HashMap", hasCycleHashMap},
	{"TwoPoints", hasCycleTwoPoints},
}

// Helper function: build a linked list from an array and a cycle entry index.
// If pos >= 0, the last node will point to the node at index pos, forming a cycle.
// If pos < 0, the list will have no cycle.
func makeCycleList(nums []int, pos int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	nodes := make([]*ListNode, len(nums))
	for i, v := range nums {
		nodes[i] = &ListNode{Val: v}
	}
	for i := 0; i < len(nums)-1; i++ {
		nodes[i].Next = nodes[i+1]
	}
	if pos >= 0 && pos < len(nums) {
		nodes[len(nums)-1].Next = nodes[pos]
	}
	return nodes[0]
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
				tc := tc
				t.Run(tc.name, func(t *testing.T) {
					head := makeCycleList(tc.nums, tc.cycleAt)
					got := algo.fn(head)
					assert.Equal(t, tc.want, got, "%s: input=%v, cycleAt=%d", algo.name, tc.nums, tc.cycleAt)
				})
			}
		})
	}
}

func BenchmarkSingleNumber(b *testing.B) {
	nums, cycleAt := []int{3, 2, 0, -4}, 1

	for _, algo := range algorithms {
		b.Run(algo.name, func(b *testing.B) {
			for b.Loop() {
				head := makeCycleList(nums, cycleAt)
				algo.fn(head)
			}
		})
	}
}
