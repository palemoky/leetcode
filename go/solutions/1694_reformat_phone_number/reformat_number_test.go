package reformat_phonenumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReformatNumber(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name   string
		number string
		want   string
	}{
		{
			name:   "Basic grouping",
			number: "1-23-45 6",
			want:   "123-456",
		},
		{
			name:   "Multiple groups",
			number: "123 4-567",
			want:   "123-45-67",
		},
		{
			name:   "Tail two digits",
			number: "123 4-5678",
			want:   "123-456-78",
		},
		{
			name:   "Long input with mixed separators",
			number: "--17-5 229 35-39475 ",
			want:   "175-229-353-94-75",
		},
		{
			name:   "Short input",
			number: "12",
			want:   "12",
		},
		{
			name:   "Single digit",
			number: "6",
			want:   "6",
		},
		{
			name:   "Only separators",
			number: "---  ",
			want:   "",
		},
		{
			name:   "Empty input",
			number: "",
			want:   "",
		},
	}

	funcsToTest := map[string]func(number string) string{
		"Basic":     reformatNumberBasic,
		"Optimized": reformatNumberOptimized,
	}

	for fnName, fn := range funcsToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					got := fn(tc.number)
					assert.Equal(t, tc.want, got, "Input: num=%s", tc.number)
				})
			}
		})
	}
}
