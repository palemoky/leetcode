package number_of_provinces

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindCircleNum(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name        string
		isConnected [][]int
		expected    int
	}{
		{
			name: "示例1：两个省份",
			isConnected: [][]int{
				{1, 1, 0},
				{1, 1, 0},
				{0, 0, 1},
			},
			expected: 2,
		},
		{
			name: "示例2：三个独立省份",
			isConnected: [][]int{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
			expected: 3,
		},
		{
			name: "单个城市",
			isConnected: [][]int{
				{1},
			},
			expected: 1,
		},
		{
			name: "所有城市相连",
			isConnected: [][]int{
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
			},
			expected: 1,
		},
		{
			name: "链式连接",
			isConnected: [][]int{
				{1, 1, 0, 0},
				{1, 1, 1, 0},
				{0, 1, 1, 1},
				{0, 0, 1, 1},
			},
			expected: 1,
		},
		{
			name: "两个独立的组",
			isConnected: [][]int{
				{1, 1, 0, 0},
				{1, 1, 0, 0},
				{0, 0, 1, 1},
				{0, 0, 1, 1},
			},
			expected: 2,
		},
	}

	funcsToTest := map[string]func([][]int) int{
		"findCircleNum": findCircleNum,
	}

	for fnName, fn := range funcsToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					result := fn(tc.isConnected)
					assert.Equal(t, tc.expected, result)
				})
			}
		})
	}
}
