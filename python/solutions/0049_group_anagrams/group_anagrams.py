class Solution:
    # Solution 1: 排序解法
    # Time: O(nlogn), Space: O(n)
    def groupAnagramsSort(self, strs: list[str]) -> list[list[str]]:
        groups = {}
        for s in strs:
            key = "".join(sorted(s))
            groups.setdefault(key, []).append(s)

        return list(groups.values())

    # Solution 2: 计数解法
    # Time: O(n), Space: O(n)
    def groupAnagramsCounter(self, strs: list[str]) -> list[list[str]]:
        groups = {}
        for s in strs:
            count = [0] * 26
            for ch in s:
                count[ord(ch) - ord("a")] += 1
            key = tuple(count)
            groups.setdefault(key, []).append(s)

        return list(groups.values())
