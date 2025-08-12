package palindrome_linked_list

import (
	"leetcode/go/solutions/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name  string
		input []int
		want  bool
	}

	testCases := []testCase{
		{
			name:  "Empty List",
			input: []int{},
			want:  true,
		},
		{
			name:  "Single Node List",
			input: []int{1},
			want:  true,
		},
		{
			name:  "Simple Palindrome with Odd Length",
			input: []int{1, 2, 1},
			want:  true,
		},
		{
			name:  "Simple Palindrome with Even Length",
			input: []int{1, 2, 2, 1},
			want:  true,
		},
		{
			name:  "Longer Palindrome",
			input: []int{1, 0, 3, 0, 1},
			want:  true,
		},
		{
			name:  "Simple Non-Palindrome",
			input: []int{1, 2, 3},
			want:  false,
		},
		{
			name:  "Non-Palindrome with Even Length",
			input: []int{1, 2, 3, 1},
			want:  false,
		},
		{
			name:  "List with all same elements",
			input: []int{5, 5, 5, 5, 5},
			want:  true,
		},
		{
			name:  "List where only ends differ",
			input: []int{1, 2, 2, 3},
			want:  false,
		},
	}

	functionsToTest := map[string]func(*utils.ListNode) bool{
		"Array":               isPalindromeArray,
		"TwoPointers&Reverse": isPalindromeTwoPointersAndReverse,
	}

	for fnName, fn := range functionsToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					head := utils.NewList(tc.input)
					got := fn(head)
					assert.Equal(t, tc.want, got)

					restoredSlice := utils.ToSlice(head)
					assert.Equal(t, tc.input, restoredSlice, "The list should be restored to its original state")
				})
			}
		})
	}
}
