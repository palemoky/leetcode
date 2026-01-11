package remove_duplicates_from_sorted_array_ii

// 解法一：快慢指针
// 核心思路：slow 左侧为重复元素的滑动窗口，通过比较 nums[slow-2] 和 nums[fast] 判断重复元素是否填满窗口
// 解题模板可查看 core/patterns/two_pointers/README.md 最多 k 个重复问题部分
// Time: O(n), Space: O(1)
func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	// slow 指向"下一个要写入的位置"
	slow := 2
	for fast := 2; fast < len(nums); fast++ {
		// 关键：比较 nums[slow-2] 和 nums[fast]
		// 如果不同，说明即使添加 nums[fast]，最多也只有2个重复
		if nums[slow-2] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
		// 如果 nums[slow-2] == nums[fast]，说明已经有2个相同的了，跳过
	}

	return slow
}
