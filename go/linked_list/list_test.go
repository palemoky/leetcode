package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestInterfaceComplianceAndToSlice tests that both list types correctly implement the interface
// and can be used by the generic toSlice function.
func TestInterfaceComplianceAndToSlice(t *testing.T) {
	testCases := []struct {
		name          string
		listToTest    List // Use the interface type here!
		expectedSlice []int
	}{
		{
			name:          "Using a SinglyList",
			listToTest:    NewSinglyList([]int{1, 1, 2, 3, 5}),
			expectedSlice: []int{1, 1, 2, 3, 5},
		},
		{
			name:          "Using a DoublyList",
			listToTest:    NewDoublyList([]int{8, 13, 21}),
			expectedSlice: []int{8, 13, 21},
		},
		{
			name:          "Using an empty SinglyList",
			listToTest:    NewSinglyList([]int{}),
			expectedSlice: []int{},
		},
		{
			name:          "Using an empty DoublyList",
			listToTest:    NewDoublyList(nil), // Test with nil input as well
			expectedSlice: []int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)
			actualSlice := toSlice(tc.listToTest)
			assert.Equal(tc.expectedSlice, actualSlice, "toSlice should work correctly for any type implementing the List interface")
		})
	}
}

// TestReverseSlice tests the simple slice reversal helper function.
func TestReverseSlice(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"Even number of elements", []int{1, 2, 3, 4}, []int{4, 3, 2, 1}},
		{"Odd number of elements", []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{"Single element", []int{1}, []int{1}},
		{"Empty slice", []int{}, []int{}},
		{"Nil slice", nil, []int{}}, // Note: Our implementation returns an empty slice for nil input
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)
			reversed := reverseSlice(tc.input)
			if tc.input == nil {
				assert.NotNil(reversed, "reverseSlice(nil) should return an empty slice, not nil")
			}
			assert.Equal(tc.expected, reversed)
		})
	}
}
