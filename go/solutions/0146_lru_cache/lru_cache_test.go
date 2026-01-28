package lru_cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// CacheInterface 定义缓存接口，用于统一测试不同实现
type CacheInterface interface {
	Get(key int) int
	Put(key int, value int)
}

func TestLRUCache(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name       string
		capacity   int
		operations []string
		args       [][]int
		expected   []any
	}{
		{
			name:       "Example 1",
			capacity:   2,
			operations: []string{"put", "put", "get", "put", "get", "put", "get", "get", "get"},
			args:       [][]int{{1, 1}, {2, 2}, {1}, {3, 3}, {2}, {4, 4}, {1}, {3}, {4}},
			expected:   []any{nil, nil, 1, nil, -1, nil, -1, 3, 4},
		},
		{
			name:       "Single capacity",
			capacity:   1,
			operations: []string{"put", "get", "put", "get", "get"},
			args:       [][]int{{1, 1}, {1}, {2, 2}, {1}, {2}},
			expected:   []any{nil, 1, nil, -1, 2},
		},
		{
			name:       "Update existing key",
			capacity:   2,
			operations: []string{"put", "put", "get", "put", "get", "get"},
			args:       [][]int{{1, 1}, {2, 2}, {1}, {1, 10}, {1}, {2}},
			expected:   []any{nil, nil, 1, nil, 10, 2},
		},
		{
			name:       "Access order matters",
			capacity:   2,
			operations: []string{"put", "put", "get", "put", "get", "get"},
			args:       [][]int{{1, 1}, {2, 2}, {1}, {3, 3}, {2}, {3}},
			expected:   []any{nil, nil, 1, nil, -1, 3},
		},
		{
			name:       "Get non-existent key",
			capacity:   2,
			operations: []string{"get"},
			args:       [][]int{{1}},
			expected:   []any{-1},
		},
		{
			name:       "Multiple updates",
			capacity:   2,
			operations: []string{"put", "put", "put", "get"},
			args:       [][]int{{1, 1}, {1, 2}, {1, 3}, {1}},
			expected:   []any{nil, nil, nil, 3},
		},
		{
			name:       "Capacity 1 eviction",
			capacity:   1,
			operations: []string{"put", "get", "put", "get", "get"},
			args:       [][]int{{1, 1}, {1}, {2, 2}, {1}, {2}},
			expected:   []any{nil, 1, nil, -1, 2},
		},
	}

	constructors := map[string]func(int) CacheInterface{
		"Manual Double Linked List": func(capacity int) CacheInterface {
			cache := Constructor(capacity)
			return &cache
		},
		"Container List": func(capacity int) CacheInterface {
			cache := Constructor1(capacity)
			return &cache
		},
	}

	for constructorName, constructor := range constructors {
		t.Run(constructorName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					cache := constructor(tc.capacity)
					for i, op := range tc.operations {
						switch op {
						case "put":
							cache.Put(tc.args[i][0], tc.args[i][1])
						case "get":
							result := cache.Get(tc.args[i][0])
							assert.Equal(t, tc.expected[i], result,
								"Operation %d: get(%d)", i, tc.args[i][0])
						}
					}
				})
			}
		})
	}
}
