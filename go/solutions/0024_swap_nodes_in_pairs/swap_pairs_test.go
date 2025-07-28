package swap_pairs

import (
	"leetcode/go/solutions/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSwapPairs(t *testing.T) {
	testCases := []struct {
		name         string
		inputList    []int
		expectedList []int
	}{
		{
			name:         "Even number of nodes",
			inputList:    []int{1, 2, 3, 4},
			expectedList: []int{2, 1, 4, 3},
		},
		{
			name:         "Odd number of nodes",
			inputList:    []int{1, 2, 3, 4, 5},
			expectedList: []int{2, 1, 4, 3, 5},
		},
		{
			name:         "Empty list",
			inputList:    []int{},
			expectedList: []int{},
		},
		{
			name:         "Single node list",
			inputList:    []int{1},
			expectedList: []int{1},
		},
		{
			name:         "Two nodes list",
			inputList:    []int{1, 2},
			expectedList: []int{2, 1},
		},
	}

	funcsToTest := map[string]func(*utils.ListNode) *utils.ListNode{
		"Iterative": swapPairsIterative,
		"Recursive": swapPairsRecursive,
	}

	for funcName, swapFunc := range funcsToTest {
		t.Run(funcName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					head := utils.CreateLinkedList(tc.inputList)
					resultHead := swapFunc(head)
					resultSlice := utils.ToSlice(resultHead)
					assert.Equal(t, tc.expectedList, resultSlice)
				})
			}
		})
	}
}
