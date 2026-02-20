package group_anagrams

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

var funcsToTest = map[string]func([]string) [][]string{
	"Sorting":  groupAnagramsSorting,
	"Counting": groupAnagramsCounting,
}

func TestGroupAnagrams(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name string
		strs []string
		want [][]string
	}{
		{"Example 1", []string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}}},
		{"Example 2", []string{""}, [][]string{{""}}},
		{"Example 3", []string{"a"}, [][]string{{"a"}}},
		{"All same", []string{"abc", "abc", "abc"}, [][]string{{"abc", "abc", "abc"}}},
		{"No anagrams", []string{"abc", "def", "ghi"}, [][]string{{"abc"}, {"def"}, {"ghi"}}},
		{"Empty strings", []string{"", ""}, [][]string{{"", ""}}},
		{"Single char anagrams", []string{"a", "b", "a"}, [][]string{{"a", "a"}, {"b"}}},
	}

	for fnName, fn := range funcsToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					got := fn(tc.strs)
					assert.Equal(t, len(tc.want), len(got), "%s: group count mismatch", fnName)
					assert.ElementsMatch(t, sortGroups(tc.want), sortGroups(got), "%s: input=%v", fnName, tc.strs)
				})
			}
		})
	}
}

// sortGroups 将每个分组内部排序，再将所有分组按首元素排序，以便比较
func sortGroups(groups [][]string) [][]string {
	result := make([][]string, len(groups))
	for i, group := range groups {
		sorted := make([]string, len(group))
		copy(sorted, group)
		sort.Strings(sorted)
		result[i] = sorted
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i][0] < result[j][0]
	})
	return result
}
