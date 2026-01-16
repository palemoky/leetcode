# 递归

对于大多数人来说，递归非常难以理解，因为人脑的限制，当我们思考递归过程时，非常容易“爆栈”。不过，我们可以通过下地下室取东西来类比递归：

1. 我们从地面开始下地下室，每到达一层都有 3 种选择：
    - 把东西打包好带上 → 前序遍历（已知当前层，未知后续层）
    - 先把东西打包好，等到返回地面时再带走 → 中序遍历（仅限 BST）
    - 在触底返回时打包带走 → 后序遍历（已知所有走过的路）
2. 触底时开始返回
3. 返回地面时，我们取回所有所需物品，递归结束。

在这个过程中，我们在每层重复的动作就是递归中的最小重复单元（这就是我们编写的递归代码），而触底返回，就是递归的终止条件。所以我们编写递归时，只要抓住了最小重复单元和终止条件，递归就迎刃而解了。

<div align="center">
  <table border="0" cellpadding="0" cellspacing="0">
    <tr>
      <td align="center" style="padding: 10px;">
        <img src="recursion_basement_analogy.png" alt="递归的地下室类比" width="60%" />
        <br/>
        <sub style="font-size: 14px;">递归的地下室类比</sub>
      </td>
      <td align="center" style="padding: 10px;">
        <pre style="text-align: left; background: var(--md-code-bg-color); padding: 15px; border-radius: 5px; font-family: monospace; border: 1px solid var(--md-default-fg-color--lightest);">
🕳️ 地面
└─ 👣 第1层：factorial(3)
   └─ 👣 第2层：factorial(2)
      └─ 👣 第3层：factorial(1)
         └─ 🎯 触底：factorial(0)
             ↓
         ┌─ ✅ 返回 1 （终止条件）
      ┌─ 🔙 返回 1 × 1 = 1
   ┌─ 🔙 返回 2 × 1 = 2
┌─ 🔙 返回 3 × 2 = 6
🌟 最终结果：6
        </pre>
        <sub style="font-size: 14px;">3! 递归调用栈示意图</sub>
      </td>
    </tr>
  </table>
</div>

常见的递归场景：

- 计算阶乘
- 前缀和数组
- 树的遍历
- 链表反转

!!! Tip

    只要递归状态从 0 开始，数组就要多开 1 个以避免越界访问，如动态规划和前缀和。

## 经典题目

- [70. 爬楼梯](https://github.com/palemoky/leetcode/blob/main/go/solutions/0070_climbing_stairs/climbing_stairs.go)
- [509. 斐波那契数](https://github.com/palemoky/leetcode/blob/main/go/solutions/0509_fibonacci_number/fibonacci_number.go)
