package subarray_sum_equals_k

// Solution 1: 暴力解法
// Time: O(n^2), Space: O(1)
func subarraySumBruteForce(nums []int, k int) int {
	count := 0
	for start := 0; start < len(nums); start++ {
		sum := 0
		for end := start; end < len(nums); end++ {
			sum += nums[end]
			if sum == k {
				count++
			}
		}
	}

	return count
}

// Solution 2: 前缀和+哈希表
// Time: O(n), Space: O(n)
func subarraySumPrefixSum(nums []int, k int) int {
	count, prefixSum := 0, 0
	prefixSumMap := map[int]int{}
	prefixSumMap[0] = 1
	for i := range len(nums) {
		prefixSum += nums[i] // 计算前缀和
		if _, ok := prefixSumMap[prefixSum-k]; ok {
			count += prefixSumMap[prefixSum-k]
		}
		prefixSumMap[prefixSum]++ // 对前缀和值的出现次数统计
	}

	return count
}
