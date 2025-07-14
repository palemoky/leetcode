package longestcommonprefix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var algorithms = []struct {
	name string
	fn   func([]string) string
}{
	{"VerticalScanning", longestCommonPrefixVerticalScanning},
	{"HorizontalScanningBuiltin", longestCommonPrefixHorizontalScanningBuiltin},
	{"HorizontalScanningByIndex", longestCommonPrefixHorizontalScanningByIndex},
}

func TestLongestCommonPrefix(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		expected string
	}{
		{name: "Basic case with common prefix", input: []string{"flower", "flow", "flight"}, expected: "fl"},
		{name: "No common prefix", input: []string{"dog", "racecar", "car"}, expected: ""},
		{name: "All strings are identical", input: []string{"hello", "hello", "hello"}, expected: "hello"},
		{name: "One string is a prefix of others", input: []string{"apple", "apply", "app"}, expected: "app"},
		{name: "Input contains an empty string", input: []string{"start", "", "station"}, expected: ""},
		{name: "Input with empty string at the beginning", input: []string{"", "b", "c"}, expected: ""},
		{name: "Single string in input", input: []string{"a"}, expected: "a"},
		{name: "Single empty string in input", input: []string{""}, expected: ""},
		{name: "Empty input slice", input: []string{}, expected: ""},
		{name: "Nil input slice", input: nil, expected: ""}, // nil 也是一种边界情况
		{name: "Long common prefix", input: []string{"interstellar", "intersection", "internal"}, expected: "inter"},
	}

	for _, algo := range algorithms {
		t.Run(algo.name, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					got := algo.fn(tc.input)
					assert.Equal(t, tc.expected, got, "Algorithm %s failed for input %v", algo.name, tc.input)
				})
			}
		})
	}
}

func generateTestStrings(count int, str1, str2 string) []string {
	strs := make([]string, count)
	for i := 0; i < count; i++ {
		if i%2 == 0 {
			strs[i] = str1
		} else {
			strs[i] = str2
		}
	}
	return strs
}

func BenchmarkLongestCommonPrefix(b *testing.B) {
	inputs := map[string][]string{
		"bestCaseInput":    {"a", "b", "c", "d", "e", "f", "g"},
		"averageCaseInput": {"flower", "flow", "flight"},
		"worstCaseInput":   generateTestStrings(100, "abcdefghijklmnopqrstuvwxyz", "abcdefghijklmnopqrstuvwxyz"),
		"longFirstInput":   generateTestStrings(100, "abcdefghijklmnopqrstuvwxyz", "abc"),
	}

	for _, input := range inputs {
		for _, algo := range algorithms {
			b.Run(algo.name, func(b *testing.B) {
				for b.Loop() {
					algo.fn(input)
				}
			})
		}
	}
}
