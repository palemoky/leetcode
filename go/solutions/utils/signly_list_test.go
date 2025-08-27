package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNewList tests the NewList function.
func TestNewList(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Normal case with multiple values",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Case with a single value",
			input:    []int{100},
			expected: []int{100},
		},
		{
			name:     "Edge case with an empty slice",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Edge case with a nil slice",
			input:    nil,
			expected: []int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			listHead := NewList(tc.input)
			actualSlice := ToSlice(listHead)
			assert.Equal(t, tc.expected, actualSlice, "The generated list should match the expected slice")
		})
	}
}

// TestToSlice tests the ToSlice function.
// While ToSlice is used in TestNewList, it's good practice to test public helper functions explicitly.
func TestToSlice(t *testing.T) {
	testCases := []struct {
		name     string
		input    *ListNode
		expected []int
	}{
		{
			name:     "Normal list",
			input:    NewList([]int{9, 8, 7}), // Use our own constructor for setup
			expected: []int{9, 8, 7},
		},
		{
			name:     "Single node list",
			input:    &ListNode{Val: 1},
			expected: []int{1},
		},
		{
			name:     "Edge case with a nil list head",
			input:    nil,
			expected: []int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := ToSlice(tc.input)
			assert.Equal(t, tc.expected, actual, "ToSlice conversion should be correct")
		})
	}
}

// TestNewCycleList tests the NewCycleList function.
// This test is more complex because we can't use ToSlice on a list with a cycle.
func TestNewCycleList(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		cyclePos int // Position where the tail's Next should point, -1 for no cycle
	}{
		{
			name:     "No cycle, normal list",
			input:    []int{3, 2, 0, -4},
			cyclePos: -1,
		},
		{
			name:     "Cycle to an intermediate node",
			input:    []int{3, 2, 0, -4},
			cyclePos: 1, // Tail (-4) should point to node with value 2
		},
		{
			name:     "Cycle to the head node",
			input:    []int{1, 2},
			cyclePos: 0, // Tail (2) should point to node with value 1
		},
		{
			name:     "Single node with a cycle to itself",
			input:    []int{1},
			cyclePos: 0,
		},
		{
			name:     "Single node with no cycle",
			input:    []int{1},
			cyclePos: -1,
		},
		{
			name:     "Edge case with empty input slice",
			input:    []int{},
			cyclePos: -1,
		},
		{
			name:     "Edge case with invalid cycle position (too large)",
			input:    []int{1, 2},
			cyclePos: 5, // Should result in no cycle
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			head := NewCycleList(tc.input, tc.cyclePos)

			// Handle the empty case first
			if len(tc.input) == 0 {
				assert.Nil(t, head, "Should return nil for empty input")
				return // End this subtest
			}

			// --- Verification ---
			// We will traverse the list and store references to each node.
			// This allows us to check both the linear part and the cycle pointer.

			nodes := make([]*ListNode, 0, len(tc.input))
			vals := make([]int, 0, len(tc.input))

			current := head
			for i := 0; i < len(tc.input); i++ {
				// Safety check to prevent panic on a malformed list (shorter than expected)
				if !assert.NotNil(t, current, "List should not terminate prematurely") {
					return
				}
				nodes = append(nodes, current)
				vals = append(vals, current.Val)
				current = current.Next
			}

			// 1. Verify the linear values are correct
			assert.Equal(t, tc.input, vals, "The linear values of the list should be correct")

			// 2. Verify the cycle connection
			tailNode := nodes[len(nodes)-1]

			// If a cycle is expected
			if tc.cyclePos >= 0 && tc.cyclePos < len(tc.input) {
				expectedCycleEntryNode := nodes[tc.cyclePos]
				// assert.Same checks if two pointers point to the exact same memory address.
				// This is perfect for verifying a cycle.
				assert.Same(t, expectedCycleEntryNode, tailNode.Next, "Tail should point to the correct cycle entry node")
			} else { // If no cycle is expected
				assert.Nil(t, tailNode.Next, "Tail should point to nil when there is no cycle")
			}
		})
	}
}
