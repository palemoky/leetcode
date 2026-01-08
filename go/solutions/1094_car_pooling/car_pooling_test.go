package car_pooling

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarPooling(t *testing.T) {
	tests := []struct {
		name     string
		trips    [][]int
		capacity int
		want     bool
	}{
		{
			name:     "Example 1 - Not enough capacity",
			trips:    [][]int{{2, 1, 5}, {3, 3, 7}},
			capacity: 4,
			want:     false,
		},
		{
			name:     "Example 2 - Enough capacity",
			trips:    [][]int{{2, 1, 5}, {3, 3, 7}},
			capacity: 5,
			want:     true,
		},
		{
			name:     "Single trip - within capacity",
			trips:    [][]int{{3, 2, 7}},
			capacity: 5,
			want:     true,
		},
		{
			name:     "Single trip - exceeds capacity",
			trips:    [][]int{{10, 1, 5}},
			capacity: 5,
			want:     false,
		},
		{
			name:     "Multiple overlapping trips",
			trips:    [][]int{{3, 2, 7}, {3, 7, 9}, {8, 3, 9}},
			capacity: 11,
			want:     true,
		},
		{
			name:     "Passengers get off before new ones board",
			trips:    [][]int{{9, 0, 1}, {3, 3, 7}},
			capacity: 4,
			want:     false,
		},
		{
			name:     "Edge case - exact capacity",
			trips:    [][]int{{2, 1, 5}, {3, 5, 7}},
			capacity: 3,
			want:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := carPooling(tt.trips, tt.capacity)
			assert.Equal(t, tt.want, got)
		})
	}
}
