package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 构造树
func buildTree(vals []any) *TreeNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}

	nodes := make([]*TreeNode, len(vals))
	for i, v := range vals {
		if v != nil {
			if val, ok := v.(int); ok {
				nodes[i] = &TreeNode{Val: val}
			}
		}
	}

	for i := range vals {
		if nodes[i] == nil {
			continue
		}

		leftIdx := 2*i + 1
		if leftIdx < len(vals) {
			nodes[i].Left = nodes[leftIdx]
		}

		rightIdx := 2*i + 2
		if rightIdx < len(vals) {
			nodes[i].Right = nodes[rightIdx]
		}
	}

	return nodes[0]
}
