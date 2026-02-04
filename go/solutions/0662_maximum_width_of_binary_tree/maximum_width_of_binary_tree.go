package maximum_width_of_binary_tree

import "leetcode/go/solutions/utils"

// Solution 1: 注意空节点也要被算入宽度
// 本题较难且频率较低，面试准备优先级较低
// Time: O(n), Space: O(n)
type pair struct {
	node  *utils.TreeNode
	index int
}

func widthOfBinaryTree(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}

	ans := 1
	q := []pair{{root, 1}}

	for q != nil {
		// 计算当前层的宽度
		ans = max(ans, q[len(q)-1].index-q[0].index+1)

		// 处理下一层
		tmp := q
		q = nil
		for _, p := range tmp {
			if p.node.Left != nil {
				q = append(q, pair{p.node.Left, p.index * 2})
			}
			if p.node.Right != nil {
				q = append(q, pair{p.node.Right, p.index*2 + 1})
			}
		}
	}

	return ans
}
