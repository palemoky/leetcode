package remove_duplicates_from_sorted_array

// 解法一：快慢指针
// 注意本题是有序数组，因此非常适合快慢指针同时完成重复检测与修改
// slow 指向"已处理的不重复元素的最后位置"，fast 用于遍历数组寻找新的不重复元素
// Time: O(n), Space: O(1)
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	slow := 0 // slow 指向当前不重复序列的最后一个元素
	for fast := 1; fast < len(nums); fast++ {
		if nums[slow] != nums[fast] { // 找到不同的元素
			slow++                  // slow 向前移动一位
			nums[slow] = nums[fast] // 将新元素写入 slow 位置
		}
		// 如果相同，只移动 fast（跳过重复元素）
	}

	return slow + 1 // 返回长度（索引+1）
}
