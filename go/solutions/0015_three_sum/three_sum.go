package three_sum

import (
	"fmt"
	"sort"
)

// 本题的关键点在于去重
// 另外，相比于两数之和，前者返回的是索引，而本题返回的是值，因此两数之和不能对数组排序处理，而本题则可以

// 解法一：暴力求解，数据规模增加会超时
// Time: O(n^3), Space: O(n)
func threeSumBruteForce(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}

	// 因为可能有重复结果，因此先排序
	sort.Ints(nums)

	existed := map[string]struct{}{}
	ans := [][]int{}
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					// 通过哈希表判断是否有重复结果，注意此处的 key 要用值，而不能是索引，因为是不同索引的值重复
					key := fmt.Sprintf("%d,%d,%d", nums[i], nums[j], nums[k])
					if _, ok := existed[key]; !ok {
						ans = append(ans, []int{nums[i], nums[j], nums[k]})
						existed[key] = struct{}{}
					}
				}
			}
		}
	}

	return ans
}

// 解法二：通过双指针将三数之和转换为两数之和
// 如要求手写排序，可通过插入或快排实现
// 本题中跳过重复的数是为了去重
// Time: O(n^2), Space: O(n)
func threeSumTwoPointers(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}

	sort.Ints(nums)

	ans := [][]int{}
	// 先固定第一个数
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue // 跳过重复的第一个数
		}

		// 用双指针让剩余两数之和与第一个数的和为0
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			// 注意此时已经排序
			if sum < 0 { // 和太小，需要右移靠近较大数
				left++
			} else if sum > 0 { // 和太大，需要左移靠近较小数
				right--
			} else { // 和为 0
				ans = append(ans, []int{nums[i], nums[left], nums[right]})

				// 跳过重复的第二个数
				for left < right && nums[left] == nums[left+1] {
					left++
				}

				// 跳过重复的第三个数
				for left < right && nums[right] == nums[right-1] {
					right--
				}

				left++
				right--
			}
		}
	}

	return ans
}
