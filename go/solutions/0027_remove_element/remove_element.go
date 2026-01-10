package remove_element

// 解法一：快慢指针
// 快指针遍历数组，遇到非 val 元素时复制到慢指针位置，慢指针前移指向下次修改位置。快指针遍历结束时，慢指针位置即为与val不等的元素的个数。
// Time: O(n), Space: O(1)
func removeElementTwoPointers(nums []int, val int) int {
	slow := 0
	for fast := range nums {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

// 解法二：双指针 + 交换
// 把所有 val 元素移到数组末尾：左指针遍历数组，遇到 val 元素时，移动到数组末尾
// Time: O(n), Space: O(1)
func removeElementSwap(nums []int, val int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		if nums[left] == val {
			// 将右边的元素移到左边
			nums[left] = nums[right]
			right--
		} else {
			left++
		}
	}

	return left
}
