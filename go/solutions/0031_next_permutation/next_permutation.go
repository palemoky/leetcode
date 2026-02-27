package next_permutation

// Solution 1:
// Time: O(n), Space: O(1)
func nextPermutation(nums []int) {
	// 从右往左找到第一个下降点 i
	i := len(nums) - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	if i >= 0 {
		// 从右往左找到第一个大于 nums[i] 的数 j
		j := len(nums) - 1
		for nums[j] <= nums[i] {
			j--
		}
		// 交换 i 和 j
		nums[i], nums[j] = nums[j], nums[i]
	}

	// 反转 i 之后的部分
	left, right := i+1, len(nums)-1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
