# 滑动窗口

滑动窗口（Sliding Window）是一种高效处理区间类问题的算法技巧，常用于字符串、数组等线性结构，适合查找满足条件的最长/最短子区间、子串等。

---

## 典型应用场景

- 最长/最短无重复子串（LeetCode 3）
- 固定/变长区间的和、乘积、计数（LeetCode 209、713）
- 子数组/子串满足某种条件的统计（LeetCode 76、438）
- 连续区间的最大/最小值

---

## 基本模式

### 固定窗口长度

适用于窗口长度已知的场景。

```go
for right := 0; right < len(nums); right++ {
    // 扩展窗口
    if right - left + 1 > k {
        left++ // 收缩窗口
    }
    // 统计/处理窗口
}
```

### 可变窗口长度

适用于窗口长度不固定，需要根据条件动态收缩。

## 优势

- 时间复杂度低，通常 $O(n)$
- 空间复杂度低，常用哈希表/计数器辅助
- 代码简洁，易于维护

## 注意事项

- 正确维护窗口边界（left/right）
- 动态更新窗口内的状态（如哈希表、计数器）
- 处理窗口收缩时的边界和条件

## 经典题目

- [LeetCode 3. 无重复字符的最长子串](https://leetcode.com/problems/longest-substring-without-repeating-characters/)
- [LeetCode 76. 最小覆盖子串](https://leetcode.com/problems/minimum-window-substring/)
- [LeetCode 438. 找到字符串中所有字母异位词](https://leetcode.com/problems/find-all-anagrams-in-a-string/)
- [LeetCode 209. 长度最小的子数组](https://leetcode.com/problems/minimum-size-subarray-sum/)
- [LeetCode 567. 字符串的排列](https://leetcode.com/problems/permutation-in-string/)
