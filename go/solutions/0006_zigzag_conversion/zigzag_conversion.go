package zigzag_conversion

import "strings"

// Solution 1:
// 本题的题目描述并不准确，更准确的名称应该是“反 N 字形变换”
// 解题步骤可以参考图解 zigzag.png（来自 https://leetcode.cn/problems/zigzag-conversion/solutions/21610/zzi-xing-bian-huan-by-jyd/）
// Time: O(n), Space: O(n)
func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}

	rows := make([]string, numRows)
	row, direction := 0, -1

	for _, c := range s {
		rows[row] += string(c)
		// 到达顶部或底部时，改变方向
		if row == 0 || row == numRows-1 {
			direction = -direction
		}
		// 因为行号可能递增也可能递减，所以要在确定方向后再更新行号
		row += direction
	}

	// 最后将所有字符串拼接返回
	return strings.Join(rows, "")
}
