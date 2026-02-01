package validate_binary_search_tree

import (
	"math"

	"leetcode/go/solutions/utils"
)

// 推荐解法：中序遍历
// 使用闭包，避免二级指针
func isValidBSTClosure(root *utils.TreeNode) bool {
	var prev *utils.TreeNode // 闭包捕获

	var inorder func(*utils.TreeNode) bool
	inorder = func(node *utils.TreeNode) bool {
		if node == nil {
			return true
		}

		if !inorder(node.Left) {
			return false
		}

		if prev != nil && node.Val <= prev.Val {
			return false
		}
		prev = node // 直接赋值

		return inorder(node.Right)
	}

	return inorder(root)
}

// 备选解法：中序遍历（二级指针）
// 核心思路：BST 的中序遍历必然是严格递增的，边遍历边判断
// 优点：空间复杂度 O(h)，可以提前终止，仍然体现 BST 核心性质
// Time: O(n), Space: O(h)
func isValidBSTPointer(root *utils.TreeNode) bool {
	var prev *utils.TreeNode // 记录中序遍历的前一个节点
	return inorder(root, &prev)
}

func inorder(node *utils.TreeNode, prev **utils.TreeNode) bool {
	if node == nil {
		return true
	}

	// 中序遍历：左 -> 根 -> 右
	if !inorder(node.Left, prev) {
		return false // 左子树不合法，提前终止
	}

	// 检查当前节点与前一个节点的大小关系（必须严格递增）
	if *prev != nil && node.Val <= (*prev).Val {
		return false // 不满足严格递增，提前终止
	}
	*prev = node // 更新前一个节点

	return inorder(node.Right, prev) // 继续检查右子树
}

// 备选解法：中序遍历（数组版）
// 优点：代码最简单，逻辑最直观
// 缺点：空间复杂度 O(n)，无法提前终止
// Time: O(n), Space: O(n)
func isValidBSTArray(root *utils.TreeNode) bool {
	nums := []int{}
	inorderTraversal(root, &nums)

	for i := 1; i < len(nums); i++ {
		// 必须是严格递增，相等也不行
		if nums[i] <= nums[i-1] {
			return false
		}
	}

	return true
}

func inorderTraversal(node *utils.TreeNode, result *[]int) {
	if node == nil {
		return
	}

	inorderTraversal(node.Left, result)
	*result = append(*result, node.Val)
	inorderTraversal(node.Right, result)
}

// 备选解法：递归约束（边界传递）
// 核心思路：每个节点的值必须在 (lower, upper) 范围内
// 优点：空间复杂度更优 O(h)，可以提前终止
// 缺点：不够直观，边界条件容易出错
// Time: O(n), Space: O(h)
func isValidBSTDFS(root *utils.TreeNode) bool {
	return isValid(root, int64(math.MinInt64), int64(math.MaxInt64))
}

func isValid(node *utils.TreeNode, lower, upper int64) bool {
	if node == nil {
		return true
	}

	val := int64(node.Val)
	// 检查当前节点的值是否在 (lower, upper) 范围内
	if val <= lower || val >= upper {
		return false
	}

	// 递归检查左右子树，并更新边界
	// 对于左子树，上界是当前节点的值
	// 对于右子树，下界是当前节点的值
	return isValid(node.Left, lower, val) &&
		isValid(node.Right, val, upper)
}
