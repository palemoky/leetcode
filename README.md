# LeetCode Solutions

[![Test](https://github.com/palemoky/leetcode/actions/workflows/leetbot.yml/badge.svg)](https://github.com/palemoky/leetcode/actions)
[![Docs](https://github.com/palemoky/leetcode/actions/workflows/docs.yml/badge.svg)](https://github.com/palemoky/leetcode/actions)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![License: CC BY 4.0](https://img.shields.io/badge/License-CC%20BY%204.0-lightgrey.svg)](https://creativecommons.org/licenses/by/4.0/)

> Computer science is no more about computers than astronomy is about telescopes. — E. W. Dijkstra

---

## Essential Algorithm Fundamentals

### Algorithm Basics You Must Master:

1. Linked List: traversal, reversal, merging, cycle detection
2. Binary Tree: level order(max/min depth, symmetric tree), preorder/inorder/postorder traversal (both recursive and iterative), inversion, lowest common ancestor, merge two binary trees
3. Sorting: quick sort, merge sort, insertion sort, selection sort, bubble sort, heap sort
4. Binary Search: on sorted arrays, binary trees
5. Bit Manipulation Techniques:
   - XOR (`^`): find numbers that appear only once or twice.
   - AND (&):
     - `n & (n-1)`: removes the lowest set bit, commonly used to check if a number is a power of 2.
     - `n & -n`: gets the lowest set bit.
   - `<<` and `>>`: efficient multiplication and division by 2.
   - Use XOR to find the unique unpaired number.
6. Math Techniques: fast exponentiation, [GCD/LCM](core/math/GCD&LCM/README.md), [elementary math simulation](core/patterns/elementary_math_simulation/README.md)

### Core Algorithm Ideas:

1. Hash table: trading space for time
2. Two pointers: fast/slow pointers, left/right pointers (collision pointers), sliding window
3. DFS: using a stack
4. BFS/backtracking: using a queue
5. Dynamic Programming (DP)

---

## Must-Know Classic LeetCode Problems by Category

1. Arrays & Strings

- [1. Two Sum](https://leetcode.cn/problems/two-sum/)
- [121. Best Time to Buy and Sell Stock](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/)
- [189. Rotate Array](https://leetcode.cn/problems/rotate-array/)
- [217. Contains Duplicate](https://leetcode.cn/problems/contains-duplicate/)
- [238. Product of Array Except Self](https://leetcode.cn/problems/product-of-array-except-self/)
- [3. Longest Substring Without Repeating Characters](https://leetcode.cn/problems/longest-substring-without-repeating-characters/)
- [125. Valid Palindrome](https://leetcode.cn/problems/valid-palindrome/)

2. Linked Lists

- [206. Reverse Linked List](https://leetcode.cn/problems/reverse-linked-list/)
- [21. Merge Two Sorted Lists](https://leetcode.cn/problems/merge-two-sorted-lists/)
- [141. Linked List Cycle](https://leetcode.cn/problems/linked-list-cycle/)
- [19. Remove Nth Node From End of List](https://leetcode.cn/problems/remove-nth-node-from-end-of-list/)
- [234. Palindrome Linked List](https://leetcode.cn/problems/palindrome-linked-list/)

3. Stacks & Queues

- [20. Valid Parentheses](https://leetcode.cn/problems/valid-parentheses/)
- [155. Min Stack](https://leetcode.cn/problems/min-stack/)
- [739. Daily Temperatures](https://leetcode.cn/problems/daily-temperatures/)
- [232. Implement Queue using Stacks](https://leetcode.cn/problems/implement-queue-using-stacks/)

4. Hash Tables

- [1. Two Sum](https://leetcode.cn/problems/two-sum/)
- [49. Group Anagrams](https://leetcode.cn/problems/group-anagrams/)
- [349. Intersection of Two Arrays](https://leetcode.cn/problems/intersection-of-two-arrays/)
- [146. LRU Cache](https://leetcode.cn/problems/lru-cache/)

5. Binary Search

- [704. Binary Search](https://leetcode.cn/problems/binary-search/)
- [35. Search Insert Position](https://leetcode.cn/problems/search-insert-position/)
- [153. Find Minimum in Rotated Sorted Array](https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array/)
- [33. Search in Rotated Sorted Array](https://leetcode.cn/problems/search-in-rotated-sorted-array/)

6. Dynamic Programming

- [70. Climbing Stairs](https://leetcode.cn/problems/climbing-stairs/)
- [53. Maximum Subarray](https://leetcode.cn/problems/maximum-subarray/)
- [322. Coin Change](https://leetcode.cn/problems/coin-change/)
- [300. Longest Increasing Subsequence](https://leetcode.cn/problems/longest-increasing-subsequence/)
- [72. Edit Distance](https://leetcode.cn/problems/edit-distance/)

7. Backtracking

- [78. Subsets](https://leetcode.cn/problems/subsets/)
- [46. Permutations](https://leetcode.cn/problems/permutations/)
- [39. Combination Sum](https://leetcode.cn/problems/combination-sum/)
- [79. Word Search](https://leetcode.cn/problems/word-search/)

8. Trees & Recursion

- [104. Maximum Depth of Binary Tree](https://leetcode.cn/problems/maximum-depth-of-binary-tree/)
- [226. Invert Binary Tree](https://leetcode.cn/problems/invert-binary-tree/)
- [101. Symmetric Tree](https://leetcode.cn/problems/symmetric-tree/)
- [102. Binary Tree Level Order Traversal](https://leetcode.cn/problems/binary-tree-level-order-traversal/)
- [235. Lowest Common Ancestor of a BST](https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-search-tree/)

9. Graphs (BFS & DFS)

- [200. Number of Islands](https://leetcode.cn/problems/number-of-islands/)
- [133. Clone Graph](https://leetcode.cn/problems/clone-graph/)
- [207. Course Schedule](https://leetcode.cn/problems/course-schedule/)
- [127. Word Ladder](https://leetcode.cn/problems/word-ladder/)

10. Greedy Algorithms

- [55. Jump Game](https://leetcode.cn/problems/jump-game/)
- [455. Assign Cookies](https://leetcode.cn/problems/assign-cookies/)
- [435. Non-overlapping Intervals](https://leetcode.cn/problems/non-overlapping-intervals/)

11. Math & Bit Manipulation

- [231. Power of Two](https://leetcode.cn/problems/power-of-two/)
- [69. Sqrt(x)](https://leetcode.cn/problems/sqrtx/)
- [371. Sum of Two Integers](https://leetcode.cn/problems/sum-of-two-integers/)

---

## Visualization Algorithm Tools

- [Algorithm Visualizer](https://algorithm-visualizer.org/)
- [VisuAlgo](https://visualgo.net/)

## Algorithm Learning Resources

- [Hello 算法](https://www.hello-algo.com/)
- [Labuladong 的算法笔记](https://labuladong.online/algo/)
- [代码随想录](https://programmercarl.com)

---

Todo

- [ ] 把 core 中的代码改用 Python 实现，聚焦与算法而非实现语言
- [ ] 英文支持

---

## 📄 License

This project is dual-licensed:

- **Code** (Go, Python, etc.): [MIT License](LICENSE)
- **Documentation & Images**: [CC BY 4.0](https://creativecommons.org/licenses/by/4.0/)

Feel free to use, modify, and share this project. Attribution is appreciated! ❤️
