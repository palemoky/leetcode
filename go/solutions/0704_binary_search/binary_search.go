package binary_search

// 解法一：迭代解法，通过不断缩小搜索窗口来确定插入位置
func searchIterative(nums []int, target int) int {
	left, right := 0, len(nums)-1

	// 这里的二分查找的核心在于每次搜索都是以 mid 为单位跳跃
	for left <= right {
		mid := left + (right-left)/2
		if target < nums[mid] {
			right = mid - 1 // 左侧区间
		} else if target > nums[mid] {
			left = mid + 1 // 右侧区间
		} else {
			return mid // 找到
		}
	}

	return -1
}

// 解法二：递归解法
func searchRecursive(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	return searchRecHelper(nums, target, 0, len(nums)-1)
}

// 注意：本题采用递归解法时，使用指针而非截取 nums，以免切片索引被打断
func searchRecHelper(nums []int, target, left, right int) int {
	if left > right {
		return -1
	}

	mid := left + (right-left)/2
	if nums[mid] == target {
		return mid
	}

	if nums[mid] < target {
		return searchRecHelper(nums, target, mid+1, right)
	}

	return searchRecHelper(nums, target, left, mid-1)
}
