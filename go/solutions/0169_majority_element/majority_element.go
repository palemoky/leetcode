package majority_element

import "sort"

// Solution 1: 排序
// 因为本题保证了多数元素的出现次数 > n/2，因此排序后中位数一定是多数元素
// Time: O(nlogn), Space: O(1)
func majorityElementSorting(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

// Solution 2: 哈希表计数
// Time: O(n), Space: O(n)
func majorityElementHashMap(nums []int) int {
	numCount := map[int]int{}
	for _, num := range nums {
		numCount[num]++
	}

	var majorityNum int
	maxCount := 0
	for num, count := range numCount {
		if count > maxCount {
			maxCount = count
			majorityNum = num
		}
	}

	return majorityNum
}

// Solution 3: Boyer-Moore 投票算法
// 核心思想：维护一个候选元素和计数器，遇到相同元素计数+1，不同元素计数-1，计数为0时更换候选元素
// 多数元素的出现次数 > n/2，因此最终的候选元素一定是多数元素
// Time: O(n), Space: O(1)
func majorityElementBoyerMoore(nums []int) int {
	candidate := nums[0]
	count := 0

	for _, num := range nums {
		if count == 0 {
			candidate = num
		}

		if num == candidate {
			count++
		} else {
			count--
		}
	}

	return candidate
}
