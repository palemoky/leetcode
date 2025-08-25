# LeetCode Solutions

> Computer science is no more about computers than astronomy is about telescopes. — E. W. Dijkstra  
> 计算机科学并不在于计算机，正如天文学并不在于望远镜。

## Contents

- [LeetCode Solutions](#leetcode-solutions)
  - [Contents](#contents)
  - [Directory Structure](#directory-structure)
  - [Essential Algorithm Fundamentals](#essential-algorithm-fundamentals)
    - [Algorithm Basics You Must Master:](#algorithm-basics-you-must-master)
    - [Core Algorithm Ideas:](#core-algorithm-ideas)
  - [Must-Know Classic LeetCode Problems by Category](#must-know-classic-leetcode-problems-by-category)
  - [Visualization Algorithm Tools](#visualization-algorithm-tools)
  - [Algorithm Learning Resources](#algorithm-learning-resources)

Welcome! This repository contains my solutions to LeetCode problems, implemented in **Go**, **Python**, and **Rust**.

Each problem is organized by language and directory. Solutions include clean code, unit tests, and concise explanations where appropriate.

## Directory Structure

- 📂 `go/` — Go solutions
- 📂 `python/` — Python solutions
- 📂 `rust/` — Rust solutions

Feel free to explore, learn, and discuss.
Contributions and suggestions are always welcome!

---

## Essential Algorithm Fundamentals

### Algorithm Basics You Must Master:

1. Linked List: traversal, reversal, merging, cycle detection
2. Binary Tree: level order, preorder/inorder/postorder traversal (both recursive and iterative), inversion
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

- [1. Two Sum](https://leetcode.com/problems/two-sum/)
- [121. Best Time to Buy and Sell Stock](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/)
- [189. Rotate Array](https://leetcode.com/problems/rotate-array/)
- [217. Contains Duplicate](https://leetcode.com/problems/contains-duplicate/)
- [238. Product of Array Except Self](https://leetcode.com/problems/product-of-array-except-self/)
- [3. Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/)
- [125. Valid Palindrome](https://leetcode.com/problems/valid-palindrome/)

2. Linked Lists

- [206. Reverse Linked List](https://leetcode.com/problems/reverse-linked-list/)
- [21. Merge Two Sorted Lists](https://leetcode.com/problems/merge-two-sorted-lists/)
- [141. Linked List Cycle](https://leetcode.com/problems/linked-list-cycle/)
- [19. Remove Nth Node From End of List](https://leetcode.com/problems/remove-nth-node-from-end-of-list/)
- [234. Palindrome Linked List](https://leetcode.com/problems/palindrome-linked-list/)

3. Stacks & Queues

- [20. Valid Parentheses](https://leetcode.com/problems/valid-parentheses/)
- [155. Min Stack](https://leetcode.com/problems/min-stack/)
- [739. Daily Temperatures](https://leetcode.com/problems/daily-temperatures/)
- [232. Implement Queue using Stacks](https://leetcode.com/problems/implement-queue-using-stacks/)

4. Hash Tables

- [1. Two Sum](https://leetcode.com/problems/two-sum/)
- [49. Group Anagrams](https://leetcode.com/problems/group-anagrams/)
- [349. Intersection of Two Arrays](https://leetcode.com/problems/intersection-of-two-arrays/)
- [146. LRU Cache](https://leetcode.com/problems/lru-cache/)

5. Binary Search

- [704. Binary Search](https://leetcode.com/problems/binary-search/)
- [35. Search Insert Position](https://leetcode.com/problems/search-insert-position/)
- [153. Find Minimum in Rotated Sorted Array](https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/)
- [33. Search in Rotated Sorted Array](https://leetcode.com/problems/search-in-rotated-sorted-array/)

6. Dynamic Programming

- [70. Climbing Stairs](https://leetcode.com/problems/climbing-stairs/)
- [53. Maximum Subarray](https://leetcode.com/problems/maximum-subarray/)
- [322. Coin Change](https://leetcode.com/problems/coin-change/)
- [300. Longest Increasing Subsequence](https://leetcode.com/problems/longest-increasing-subsequence/)
- [72. Edit Distance](https://leetcode.com/problems/edit-distance/)

7. Backtracking

- [78. Subsets](https://leetcode.com/problems/subsets/)
- [46. Permutations](https://leetcode.com/problems/permutations/)
- [39. Combination Sum](https://leetcode.com/problems/combination-sum/)
- [79. Word Search](https://leetcode.com/problems/word-search/)

8. Trees & Recursion

- [104. Maximum Depth of Binary Tree](https://leetcode.com/problems/maximum-depth-of-binary-tree/)
- [226. Invert Binary Tree](https://leetcode.com/problems/invert-binary-tree/)
- [101. Symmetric Tree](https://leetcode.com/problems/symmetric-tree/)
- [102. Binary Tree Level Order Traversal](https://leetcode.com/problems/binary-tree-level-order-traversal/)
- [235. Lowest Common Ancestor of a BST](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/)

9. Graphs (BFS & DFS)

- [200. Number of Islands](https://leetcode.com/problems/number-of-islands/)
- [133. Clone Graph](https://leetcode.com/problems/clone-graph/)
- [207. Course Schedule](https://leetcode.com/problems/course-schedule/)
- [127. Word Ladder](https://leetcode.com/problems/word-ladder/)

10. Greedy Algorithms

- [55. Jump Game](https://leetcode.com/problems/jump-game/)
- [455. Assign Cookies](https://leetcode.com/problems/assign-cookies/)
- [435. Non-overlapping Intervals](https://leetcode.com/problems/non-overlapping-intervals/)

11. Math & Bit Manipulation

- [231. Power of Two](https://leetcode.com/problems/power-of-two/)
- [69. Sqrt(x)](https://leetcode.com/problems/sqrtx/)
- [371. Sum of Two Integers](https://leetcode.com/problems/sum-of-two-integers/)

---

## Visualization Algorithm Tools

- [Algorithm Visualizer](https://algorithm-visualizer.org/)
- [VisuAlgo](https://visualgo.net/)

## Algorithm Learning Resources

- [Hello 算法](https://www.hello-algo.com/)
- [Labuladong 的算法笔记](https://labuladong.online/algo/)
- [代码随想录](https://programmercarl.com)
