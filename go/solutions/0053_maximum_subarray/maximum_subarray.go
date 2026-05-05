package maximum_subarray

// Solution 1: 动态规划
// Time: O(n), Space: O(1)
// 状态转移方程：dp[i] = max(nums[i], dp[i-1] + nums[i])
// 含义：
// dp[i]：以 i 结尾的最大子数组和
// 自己重新开始（nums[i]）
// 接上之前的（dp[i-1] + nums[i]）
func maxSubArray(nums []int) int {
	sum, ans := 0, nums[0]
	for _, num := range nums {
		if sum > 0 { // sum 对结果有增益效果
			sum += num
		} else { // sum 对结果是减损效果，需要舍弃并更新为当前遍历数字
			sum = num
		}

		// 以上代码也可以进一步压缩为以下代码，但更难理解思路
		// sum = max(num, sum+num)

		ans = max(ans, sum) // 取最大值
	}

	return ans
}
