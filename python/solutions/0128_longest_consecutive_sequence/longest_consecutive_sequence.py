class Solution:
    # Solution 1:
    # Time: O(n), Space: O(n)
    def longestConsecutive(self, nums: list[int]) -> int:
        has = set(nums)

        ans = 0
        for num in nums:
            if num - 1 in has:
                continue

            next_num = num + 1
            while next_num in has:
                next_num += 1

            ans = max(ans, next_num - num)

        return ans
