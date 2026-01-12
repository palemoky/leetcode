package move_zeroes

// 解法一：快慢指针
// 注意：不能使用对撞指针直接首尾交换，因为题目要求保持非零元素的相对顺序
// 只能通过快慢指针不断前进，遇到非零元素时与 slow 位置交换
// Time: O(n), Space: O(1)
func moveZeroes(nums []int) {
	// 两者都从 0 开始，可以正确处理所有情况（前导非零、前导零、单元素等）
	slow := 0                                 // 指向下一个非零元素应该放的位置
	for fast := 0; fast < len(nums); fast++ { // 遍历所有元素
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
	}
}
