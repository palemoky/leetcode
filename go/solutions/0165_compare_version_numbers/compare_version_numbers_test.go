package compare_version_numbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var funcsToTest = map[string]func(string, string) int{
	"Simulation": compareVersion,
}

func TestCompareVersion(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		version1 string
		version2 string
		want     int
	}{
		{"Example 1", "1.2", "1.10", -1},
		{"Example 2", "1.01", "1.001", 0},
		{"Example 3", "1.0", "1.0.0.0", 0},
		{"Equal", "1.0.0", "1.0.0", 0},
		{"V1 greater", "2.0", "1.9.9", 1},
		{"V2 greater", "0.1", "1.0", -1},
		{"Leading zeros", "1.005", "1.5", 0},
		{"Multi segments", "1.2.3.4.5", "1.2.3.4.5", 0},
	}

	for fnName, fn := range funcsToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					assert.Equal(t, tc.want, fn(tc.version1, tc.version2))
				})
			}
		})
	}
}
