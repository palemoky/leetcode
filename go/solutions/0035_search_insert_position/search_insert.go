package search_insert_position

// 注意两种解法的差一和遍历边界条件

// 解法一：标准二分查找，查找目标值或插入位置
// Time: O(log n), Space: O(1)
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		// mid 计算方式为向下取整，偶数长度区间时指向中间偏左的元素，避免所有情况下的越界访问
		mid := left + (right-left)/2
		if target > nums[mid] {
			left = mid + 1
		} else if target < nums[mid] {
			right = mid - 1
		} else {
			return mid // 找到目标的插入位置
		}
	}

	return left // 未找到目标时，left 会停在第一个 >= target 的位置，也就是 target 应该插入的位置
}

// 解法二：lowerBound 实现，即在有序数组中查找第一个大于等于目标值的位置
// Time: O(log n), Space: O(1)
func searchInsertLowerBound(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left // 返回第一个 >= target 的位置
}
