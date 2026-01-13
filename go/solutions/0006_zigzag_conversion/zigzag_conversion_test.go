package zigzag_conversion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name    string
		s       string
		numRows int
		want    string
	}{
		{
			name:    "Example 1",
			s:       "PAYPALISHIRING",
			numRows: 3,
			want:    "PAHNAPLSIIGYIR",
		},
		{
			name:    "Example 2",
			s:       "PAYPALISHIRING",
			numRows: 4,
			want:    "PINALSIGYAHRPI",
		},
		{
			name:    "Example 3",
			s:       "A",
			numRows: 1,
			want:    "A",
		},
		{
			name:    "Single row",
			s:       "ABCDEFGH",
			numRows: 1,
			want:    "ABCDEFGH",
		},
		{
			name:    "Two rows",
			s:       "ABCDE",
			numRows: 2,
			want:    "ACEBD",
		},
		{
			name:    "More rows than characters",
			s:       "ABC",
			numRows: 5,
			want:    "ABC",
		},
		{
			name:    "Empty string",
			s:       "",
			numRows: 3,
			want:    "",
		},
		{
			name:    "Single character",
			s:       "X",
			numRows: 3,
			want:    "X",
		},
		{
			name:    "Five rows",
			s:       "ABCDEFGHIJKLMNOP",
			numRows: 5,
			want:    "AIBHJPCGKODFLNEM",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			result := convert(tc.s, tc.numRows)
			assert.Equal(t, tc.want, result)
		})
	}
}
