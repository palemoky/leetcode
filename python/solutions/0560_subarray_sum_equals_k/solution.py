# Solution 1: 这种连续区间问题非常适合使用前缀和来求解
# Time: O(n), Space: O(n)
def subarraySum(nums: list[int], k: int) -> int:
    count, prefix_sum = 0, 0
    prefix_sum_map = {0: 1}
    for num in nums:
        prefix_sum += num
        if prefix_sum - k in prefix_sum_map:
            count += prefix_sum_map[prefix_sum - k]
        prefix_sum_map[prefix_sum] = prefix_sum_map.get(prefix_sum, 0) + 1

    return count
