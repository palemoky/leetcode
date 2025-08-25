package utils

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const (
	WHITE = 0
	BLACK = 1
)

type ColorNode struct {
	Color int
	Node  *TreeNode
}

func BuildTree(vals []any) *TreeNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}

	nodes := make([]*TreeNode, len(vals))
	for i, v := range vals {
		if v == nil {
			nodes[i] = nil
		} else {
			nodes[i] = &TreeNode{Val: v.(int)}
		}
	}

	for i := range vals {
		if nodes[i] == nil {
			continue
		}

		leftIdx := 2*i + 1
		rightIdx := 2*i + 2
		if leftIdx < len(vals) {
			nodes[i].Left = nodes[leftIdx]
		}

		if rightIdx < len(vals) {
			nodes[i].Right = nodes[rightIdx]
		}
	}

	return nodes[0]
}
