package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToSlice(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		list     List
		expected []int
	}{
		{
			name:     "Using a SinglyList",
			list:     NewSinglyList([]int{1, 1, 2, 3, 5}),
			expected: []int{1, 1, 2, 3, 5},
		},
		{
			name:     "Using a DoublyList",
			list:     NewDoublyList([]int{8, 13, 21}),
			expected: []int{8, 13, 21},
		},
		{
			name:     "Using an empty SinglyList",
			list:     NewSinglyList([]int{}),
			expected: []int{},
		},
		{
			name:     "Using an empty DoublyList",
			list:     NewDoublyList(nil), // Test with nil input as well
			expected: []int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			result := toSlice(tc.list)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestToSliceReverse(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		list     *DoublyList
		expected []int
	}{
		{
			name:     "Nil list",
			list:     nil,
			expected: []int{},
		},
		{
			name:     "Empty list",
			list:     NewDoublyList(nil),
			expected: []int{},
		},
		{
			name:     "Single node list",
			list:     NewDoublyList([]int{1}),
			expected: []int{1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			result := toSliceReverse(tc.list)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestReverseSlice(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"Empty slice", []int{}, []int{}},
		{"Single element", []int{1}, []int{1}},
		{"Multiple elements", []int{1, 2, 3, 4}, []int{4, 3, 2, 1}},
		{"Two elements", []int{1, 2}, []int{2, 1}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			result := reverseSlice(tc.input)
			assert.Equal(t, tc.expected, result)

			// Ensure original slice is not modified
			if len(tc.input) > 1 {
				assert.NotEqual(t, tc.input, result)
			}
		})
	}
}
