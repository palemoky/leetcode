package two_sum_ii_input_array_is_sorted

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name    string
		numbers []int
		target  int
		want    []int
	}{
		// --- 1. 基本情况 ---
		{
			name:    "Two elements - found",
			numbers: []int{1, 2},
			target:  3,
			want:    []int{1, 2},
		},
		{
			name:    "Two elements - first element",
			numbers: []int{1, 3},
			target:  4,
			want:    []int{1, 2},
		},

		// --- 2. 典型情况 ---
		{
			name:    "Example 1: target in middle",
			numbers: []int{2, 7, 11, 15},
			target:  9,
			want:    []int{1, 2},
		},
		{
			name:    "Example 2: target at end",
			numbers: []int{2, 3, 4},
			target:  6,
			want:    []int{1, 3},
		},
		{
			name:    "Example 3: negative numbers",
			numbers: []int{-1, 0},
			target:  -1,
			want:    []int{1, 2},
		},

		// --- 3. 边界情况 ---
		{
			name:    "Target at first and last",
			numbers: []int{1, 2, 3, 4, 5},
			target:  6,
			want:    []int{1, 5},
		},
		{
			name:    "Target at adjacent elements",
			numbers: []int{1, 2, 3, 4, 5},
			target:  3,
			want:    []int{1, 2},
		},
		{
			name:    "Large numbers",
			numbers: []int{1, 1000, 2000, 3000},
			target:  3001,
			want:    []int{1, 4},
		},

		// --- 4. 负数和零 ---
		{
			name:    "All negative numbers",
			numbers: []int{-5, -3, -1, 0},
			target:  -4,
			want:    []int{2, 3},
		},
		{
			name:    "Mix of negative and positive",
			numbers: []int{-10, -5, 0, 5, 10},
			target:  0,
			want:    []int{1, 5},
		},
		{
			name:    "Zero sum",
			numbers: []int{-3, -1, 0, 1, 3},
			target:  0,
			want:    []int{1, 5}, // -3 + 3 = 0
		},

		// --- 5. 重复元素 ---
		{
			name:    "Duplicates - use same value twice",
			numbers: []int{1, 2, 2, 3},
			target:  4,
			want:    []int{1, 4}, // 1 + 3 = 4
		},
		{
			name:    "All same elements",
			numbers: []int{5, 5, 5, 5},
			target:  10,
			want:    []int{1, 4}, // 5 + 5 = 10 (first and last)
		},
	}

	funcsToTest := map[string]func([]int, int) []int{
		"BruteForce":  twoSumBruteForce,
		"TwoPointers": twoSumTwoPointers,
	}

	for fnName, fn := range funcsToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					got := fn(tc.numbers, tc.target)

					// 验证答案的正确性（可能有多个正确答案）
					assert.Len(t, got, 2, "should return exactly 2 indices")
					assert.True(t, got[0] >= 1 && got[0] <= len(tc.numbers), "first index out of range")
					assert.True(t, got[1] >= 1 && got[1] <= len(tc.numbers), "second index out of range")
					assert.True(t, got[0] < got[1], "indices should be in ascending order")

					// 验证和是否等于target（索引从1开始，所以要-1）
					sum := tc.numbers[got[0]-1] + tc.numbers[got[1]-1]
					assert.Equal(t, tc.target, sum, "twoSum(%v, %d) = %v, but %d + %d = %d != %d",
						tc.numbers, tc.target, got,
						tc.numbers[got[0]-1], tc.numbers[got[1]-1], sum, tc.target)
				})
			}
		})
	}
}
