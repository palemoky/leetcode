package two_sum_ii_input_array_is_sorted

// 解法一：暴力解法（双重循环）
// Time: O(n^2), Space: O(1)
func twoSumBruteForce(numbers []int, target int) []int {
	for i := 0; i < len(numbers)-1; i++ {
		for j := i + 1; j < len(numbers); j++ {
			sum := numbers[i] + numbers[j]

			if sum == target {
				return []int{i + 1, j + 1} // 题目要求索引从1开始
			}

			if sum > target {
				break
			}
		}
	}

	return []int{}
}

// 解法：对撞双指针
// 利用数组已排序的特性，从两端向中间收缩
// Time: O(n), Space: O(1)
func twoSumTwoPointers(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1

	for left < right {
		sum := numbers[left] + numbers[right]

		if sum == target {
			return []int{left + 1, right + 1} // 题目要求索引从1开始
		} else if sum < target {
			left++ // 和太小，需要更大的数
		} else {
			right-- // 和太大，需要更小的数
		}
	}

	return []int{}
}
