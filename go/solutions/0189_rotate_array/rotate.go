package rotate_array

// 解法一：每次将末尾的元素暂存到额外变量，然后依次右移剩余数组
// Time: O(n^2), Space: O(1)
func rotateShift(nums []int, k int) {
	n := len(nums)
	for range k {
		lastElement := nums[n-1]
		for i := n - 1; i >= 1; i-- {
			nums[i] = nums[i-1]
		}
		nums[0] = lastElement
	}
}

// 解法二：使用切片处理并将结果 copy 给 nums
// Time: O(n), Space: O(n)
func rotateSlice(nums []int, k int) {
	n := len(nums)
	k %= n // 兼容 k > n
	// 因为 append 会创建新切片，所以需要使用 copy
	copy(nums, append(nums[n-k:], nums[:n-k]...))
}

// 解法三：反转数组
// 核心思想：负负得正，可查看目录下的 reverse.png 图解
// Time: O(n), Space: O(1)
func rotateReverse(nums []int, k int) {
	n := len(nums)
	k %= n

	reverse(nums)     // 反转整个数组
	reverse(nums[:k]) // 反转前 k 个
	reverse(nums[k:]) // 反转后 n-k 个
}

// 反转数组也可以用 slices.Reverse 替代
func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
