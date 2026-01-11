package move_zeroes

// 解法一：快慢指针
// 注意：不能使用对撞指针直接首尾交换，因为题目要求保持非零元素的相对顺序
// 只能通过快慢指针不断前进，遇到非零元素时与 slow 位置交换
// Time: O(n), Space: O(1)
func moveZeroes(nums []int) {
	slow := 0
	// 因为 len(nums) 范围是 [1, 10^4]，所以 fast 要从 0 开始，不能从 1 开始
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
	}
}
