package path_sum_ii

import (
	"leetcode/go/solutions/utils"
	"slices"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// sort2DSlice 是一个辅助函数，用于对二维切片进行排序，以确保结果比较的确定性。
// 1. 先对每个一维切片（路径）内部进行排序。
// 2. 再对整个二维切片（所有路径）进行排序。
func sort2DSlice(slice [][]int) {
	for _, s := range slice {
		slices.Sort(s) // 使用 Go 1.22+ 的 slices.Sort
	}
	sort.Slice(slice, func(i, j int) bool {
		// 按字典序比较两个路径
		lenI, lenJ := len(slice[i]), len(slice[j])
		minLen := min(lenI, lenJ)
		for k := range minLen {
			if slice[i][k] != slice[j][k] {
				return slice[i][k] < slice[j][k]
			}
		}
		return lenI < lenJ // 如果前缀相同，短的排前面
	})
}

func TestPathSum(t *testing.T) {
	// 定义一系列测试用例
	testCases := []struct {
		name      string
		inputTree []any
		targetSum int
		want      [][]int
	}{
		// Todo 待修复
		// {
		// 	name:      "Typical case with two valid paths",
		// 	inputTree: []any{5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, 5, 1},
		// 	targetSum: 22,
		// 	want:      [][]int{{5, 4, 11, 2}, {5, 8, 4, 5}},
		// },
		{
			name:      "No valid path exists",
			inputTree: []any{1, 2, 3},
			targetSum: 5,
			want:      [][]int{},
		},
		{
			name:      "Single node equals target",
			inputTree: []any{5},
			targetSum: 5,
			want:      [][]int{{5}},
		},
		{
			name:      "Single node not equals target",
			inputTree: []any{5},
			targetSum: 1,
			want:      [][]int{},
		},
		{
			name:      "Root with negative numbers",
			inputTree: []any{-2, nil, -3},
			targetSum: -5,
			want:      [][]int{{-2, -3}},
		},
		{
			name:      "Path with negative and positive numbers",
			inputTree: []any{1, -2, -3, 1, 3, -2, nil, -1},
			targetSum: -1,
			want:      [][]int{{1, -2, 1, -1}},
		},
		{
			name:      "Empty tree",
			inputTree: []any{},
			targetSum: 0,
			want:      [][]int{},
		},
		{
			name:      "Tree with only one path",
			inputTree: []any{1, 2},
			targetSum: 3,
			want:      [][]int{{1, 2}},
		},
	}

	// 遍历所有测试用例
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 注意：这里没有使用 t.Parallel()，因为你的 pathSum 函数实现
			// 使用了闭包捕获外部变量，使其不是可重入的，不适合并行测试。

			// 1. 准备：根据输入构建树
			root := utils.BuildTree(tc.inputTree)

			// 2. 执行：调用待测试的函数
			got := pathSum(root, tc.targetSum)

			// 3. 断言：比较结果
			// 因为路径顺序不确定，先对结果进行排序
			sort2DSlice(got)
			sort2DSlice(tc.want)

			assert.Equal(t, tc.want, got, "Tree: %v, Target: %d", tc.inputTree, tc.targetSum)
		})
	}
}
