package multiply_strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiply(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name string
		num1 string
		num2 string
		want string
	}{
		{
			name: "Single digit multiply",
			num1: "2",
			num2: "3",
			want: "6",
		},
		{
			name: "Multi-digit multiply",
			num1: "123",
			num2: "456",
			want: "56088",
		},
		{
			name: "one num is zero",
			num1: "123",
			num2: "0",
			want: "0",
		},
		{
			name: "both num is zero",
			num1: "0",
			num2: "0",
			want: "0",
		},
		{
			name: "one num is one",
			num1: "123",
			num2: "1",
			want: "123",
		},
		{
			name: "num1 is empty string",
			num1: "",
			num2: "123",
			want: "0",
		},
		{
			name: "num2 is empty string",
			num1: "123",
			num2: "",
			want: "0",
		},
		{
			name: "both empty string",
			num1: "",
			num2: "",
			want: "",
		},
		{
			name: "num1 has leading zeros",
			num1: "0000123",
			num2: "456",
			want: "56088",
		},
		{
			name: "num2 has leading zeros",
			num1: "123",
			num2: "0000456",
			want: "56088",
		},
		{
			name: "both have leading zeros",
			num1: "0000123",
			num2: "0000456",
			want: "56088",
		},
		{
			name: "very large numbers",
			num1: "9876543210123456789",
			num2: "1234567890987654321",
			want: "12193263121170553265523548251112635269",
		},
		{
			name: "long single digit numbers",
			num1: "99999999999999999999",
			num2: "9",
			want: "899999999999999999991",
		},
		{
			name: "both are long numbers",
			num1: "123456789012345678901234567890",
			num2: "987654321098765432109876543210",
			want: "121932631137021795226185032733622923332237463801111263526900",
		},
	}

	funcsToTest := map[string]func(num1 string, num2 string) string{
		"multiply": multiply,
	}

	for fnName, fn := range funcsToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					got := fn(tc.num1, tc.num2)
					assert.Equal(t, tc.want, got, "input: %s * %s", tc.num1, tc.num2)
				})
			}
		})
	}
}
