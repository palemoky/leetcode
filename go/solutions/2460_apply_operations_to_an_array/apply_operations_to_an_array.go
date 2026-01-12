package apply_operations_to_an_array

// Solution 1: 快慢指针计算+滑动窗口交换零
// Time: O(n), Space: O(1)
func applyOperationsWithSlidingWindow(nums []int) []int {
	slow := 0
	for fast := 1; fast < len(nums); fast++ {
		if nums[slow] == nums[fast] {
			nums[slow] *= 2
			nums[fast] = 0
		}

		slow++
	}

	// 排序时使用双指针交换，可将 [left, right] 看作滑动窗口，左边 left 的值为 0，右边 right 的值为非 0，区间内都是 0
	left, right := 0, 0
	for right < len(nums) {
		// 找到第一个 0
		for left < len(nums) && nums[left] != 0 {
			left++
		}

		// right 从 left 开始找第一个非 0
		right = left
		for right < len(nums) && nums[right] == 0 {
			right++
		}

		// 如果 right 没有越界，交换
		if right < len(nums) {
			nums[left], nums[right] = nums[right], nums[left]
		} else {
			break
		}
	}

	return nums
}

// Solution 2: Solution 1 移动零部分用 283 题的优化版本
// Time: O(n), Space: O(1)
func applyOperations(nums []int) []int {
	slow := 0
	for fast := 1; fast < len(nums); fast++ {
		if nums[slow] == nums[fast] {
			nums[slow] *= 2
			nums[fast] = 0
		}

		slow++
	}

	// 用 283 题的方法移动 0
	left, right := 0, 0
	for right < len(nums) {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}

	return nums
}

// Solution 3: 单次遍历（一边计算一边移动非零元素）
// Time: O(n), Space: O(1)
func applyOperationsOnePass(nums []int) []int {
	// fast 计算，slow 移动非零元素
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if fast < len(nums)-1 && nums[fast] == nums[fast+1] {
			nums[fast] *= 2
			nums[fast+1] = 0
		}
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
	}

	return nums
}
