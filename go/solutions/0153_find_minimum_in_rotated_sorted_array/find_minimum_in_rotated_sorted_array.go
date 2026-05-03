package find_minimum_in_rotated_sorted_array

// Solution 1: 二分法
// Time: O(logn), Space: O(1)
func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < nums[right] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return nums[left]
}
