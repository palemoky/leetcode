package power_of_four

import "math"

func isPowerOfFourIterative(n int) bool {
	if n <= 0 {
		return false
	}

	for n%4 == 0 {
		n /= 4
	}

	return n == 1
}

func isPowerOfFourLog(n int) bool {
	if n <= 0 {
		return false
	}

	k := math.Log10(float64(n)) / math.Log10(4)
	return math.Abs(k-math.Round(k)) < 1e-10
}
