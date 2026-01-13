package rearrange_spaces_between_words

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRearrangeSpaces(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name string
		text string
		want string
	}{
		{
			"Example 1",
			"  this   is  a sentence ",
			"this   is   a   sentence",
		},
		{
			"Example 2",
			" practice   makes   perfect",
			"practice   makes   perfect ",
		},
		{
			"Single word with spaces",
			"  hello  ",
			"hello    ",
		},
		{
			"Two words",
			"a  b",
			"a  b", // 2 spaces / 1 gap = 2 spaces per gap
		},
		{
			"No spaces",
			"hello",
			"hello",
		},
		{
			"Multiple consecutive spaces",
			"a    b    c",
			"a    b    c", // 8 spaces / 2 gaps = 4 spaces per gap
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := rearrangeSpaces(tc.text)
			assert.Equal(t, tc.want, got)
		})
	}
}
