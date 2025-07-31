package time_needed_to_buy_tickets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var funcsToTest = map[string]func([]int, int) int{
	"Queue":   timeRequiredToBuyQueue,
	"Pointer": timeRequiredToBuyPointer,
	"Math":    timeRequiredToBuyMath,
}

func TestTimeRequiredToBuy(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name    string
		tickets []int
		k       int
		want    int
	}{
		{"Target at end", []int{1, 2, 3}, 2, 6},
		{"All same", []int{2, 2, 2, 2}, 1, 6},
		{"Target needs one", []int{3, 1, 2}, 1, 2},
		{"Target most tickets", []int{1, 2, 5, 1}, 2, 9},
		{"Only target", []int{7}, 0, 7},
		{"Empty", []int{}, 0, 0},
		{"All zero", []int{0, 0, 0}, 1, 0},
	}

	for fnName, fn := range funcsToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					ticketsCopy := make([]int, len(tc.tickets))
					copy(ticketsCopy, tc.tickets)
					got := fn(ticketsCopy, tc.k)
					assert.Equal(t, tc.want, got, "%s: input=%v", fnName, tc.tickets, tc.k)
				})
			}
		})
	}
}

func BenchmarkTimeRequiredToBuy(b *testing.B) {
	tickets, k := []int{5, 1, 1, 1}, 0
	for fnName, fn := range funcsToTest {
		b.Run(fnName, func(b *testing.B) {
			for b.Loop() {
				t := make([]int, len(tickets))
				copy(t, tickets)
				fn(t, k)
			}
		})
	}
}
