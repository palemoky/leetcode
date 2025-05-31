package leetcode

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		name   string 
		nums   []int
		target int
		want   []int
	}{
		{"Example 1", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"Example 2", []int{3, 2, 4}, 6, []int{1, 2}},
		{"Example 3", []int{3, 3}, 6, []int{0, 1}},
	}

	// 测试暴力解法
	t.Run("BruteForce", func(t *testing.T) {
		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				got := twoSumBruteForce(tc.nums, tc.target) 
				if !reflect.DeepEqual(got, tc.want) && !reflect.DeepEqual(got, []int{tc.want[1], tc.want[0]}) { // 简单处理顺序问题
					t.Errorf("twoSumBruteForce() = %v; want %v", got, tc.want)
				}
			})
		}
	})

	// 测试哈希表解法
	t.Run("HashMap", func(t *testing.T) {
		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				got := twoSumHashMap(tc.nums, tc.target) 
				if !reflect.DeepEqual(got, tc.want) && !reflect.DeepEqual(got, []int{tc.want[1], tc.want[0]}) {
					t.Errorf("twoSumHashMap() = %v; want %v", got, tc.want)
				}
			})
		}
	})
}
