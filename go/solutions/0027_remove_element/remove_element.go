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

// nums = [2, 3, 2], val = 2
// Init: left=0, right=2
//
//	[2, 3, 2]
//	 ↑     ↑
//
// Step 1: nums[0]==2，swap
//
//	[2, 3, 2] → nums[0]=nums[2]
//	 ↑     ↑
//	right--
//	[2, 3, 2]
//	 ↑  ↑
//
// Step 2: nums[0]==2，swap
//
//	[3, 3, 2] → nums[0]=nums[1]
//	 ↑  ↑
//	right--
//	[3, 3, 2]
//	 ↑↑
//
// Step 3: nums[0]==3，不是val
//
//	left++
//	[3, 3, 2]
//	  ↑
//	left > right，done
//

// Time: O(n), Space: O(1)
func removeElementSwap(nums []int, val int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		if nums[left] == val {
			// 因为题目中并不要求 k 个元素之外的元素，因此只需将右边的元素移到左边即可
			nums[left] = nums[right]
			// 注意此时不能 left++，因为右边交换过来的元素也可能是 val
			right--
		} else {
			left++
		}
	}

	return left
}
