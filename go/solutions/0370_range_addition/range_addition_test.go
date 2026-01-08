package range_addition

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetModifiedArray(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		length   int
		updates  [][]int
		expected []int
	}{
		{
			name:   "example 1: basic updates",
			length: 5,
			updates: [][]int{
				{1, 3, 2},
				{2, 4, 3},
				{0, 2, -2},
			},
			expected: []int{-2, 0, 3, 5, 3},
		},
		{
			name:   "example 2: single update",
			length: 10,
			updates: [][]int{
				{2, 4, 6},
			},
			expected: []int{0, 0, 6, 6, 6, 0, 0, 0, 0, 0},
		},
		{
			name:   "example 3: multiple overlapping updates",
			length: 5,
			updates: [][]int{
				{1, 3, 2},
				{2, 4, 3},
			},
			expected: []int{0, 2, 5, 5, 3},
		},
		{
			name:     "empty updates",
			length:   5,
			updates:  [][]int{},
			expected: []int{0, 0, 0, 0, 0},
		},
		{
			name:   "single element array",
			length: 1,
			updates: [][]int{
				{0, 0, 5},
			},
			expected: []int{5},
		},
		{
			name:   "update entire array",
			length: 3,
			updates: [][]int{
				{0, 2, 10},
			},
			expected: []int{10, 10, 10},
		},
		{
			name:   "negative increments",
			length: 5,
			updates: [][]int{
				{0, 4, 5},
				{1, 3, -3},
			},
			expected: []int{5, 2, 2, 2, 5},
		},
		{
			name:   "consecutive non-overlapping updates",
			length: 6,
			updates: [][]int{
				{0, 1, 1},
				{2, 3, 2},
				{4, 5, 3},
			},
			expected: []int{1, 1, 2, 2, 3, 3},
		},
		{
			name:   "update first element only",
			length: 5,
			updates: [][]int{
				{0, 0, 7},
			},
			expected: []int{7, 0, 0, 0, 0},
		},
		{
			name:   "update last element only",
			length: 5,
			updates: [][]int{
				{4, 4, 7},
			},
			expected: []int{0, 0, 0, 0, 7},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := getModifiedArray(tt.length, tt.updates)
			assert.Equal(t, tt.expected, result)
		})
	}
}
