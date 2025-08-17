package first_bad_version

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstBadVersion(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		n      int
		badVer int
		want   int
	}{
		{
			name:   "Bad version is first",
			n:      5,
			badVer: 1,
			want:   1,
		},
		{
			name:   "Bad version is last",
			n:      5,
			badVer: 5,
			want:   5,
		},
		{
			name:   "Bad version is middle",
			n:      5,
			badVer: 3,
			want:   3,
		},
		{
			name:   "Single version",
			n:      1,
			badVer: 1,
			want:   1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			bad = tc.badVer // 设置全局变量
			got := firstBadVersion(tc.n)
			assert.Equal(t, tc.want, got, "n=%d, bad=%d", tc.n, tc.badVer)
		})
	}
}
