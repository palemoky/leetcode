package range_sum_query_immutable

type NumArray struct {
	preSum []int
}

// 解法一：前缀和
// Time: O(n), Space: O(n)
func Constructor(nums []int) NumArray {
	preSum := make([]int, len(nums)+1)
	for i := range nums {
		preSum[i+1] = preSum[i] + nums[i]
	}

	return NumArray{preSum}
}

// Time: O(1), Space: O(1)
func (this *NumArray) SumRange(left int, right int) int {
	return this.preSum[right+1] - this.preSum[left]
}
