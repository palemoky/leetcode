package rotate_array

// 本题如果有切片返回值，一行代码就可解 return append(nums[len(nums)-k:], nums[:len(nums)-k]...)

// 解法一：每次将末尾的元素暂存到额外变量，然后依次右移剩余数组
// Time: O(n^2), Space: O(1)
func rotateBruteForce(nums []int, k int) {
	n := len(nums)
	for range k {
		tmp := nums[n-1]
		for i := n - 1; i >= 1; i-- {
			nums[i] = nums[i-1]
		}
		nums[0] = tmp
	}
}

// 解法二（推荐）：使用切片处理并将结果 copy 给 nums
func rotateSlice(nums []int, k int) {
	n := len(nums)
	k %= n // 兼容 k > n
	copy(nums, append(nums[n-k:], nums[:n-k]...))
}
