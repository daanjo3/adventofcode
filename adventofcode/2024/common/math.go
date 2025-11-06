package common

import "math"

func Abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func IntPow(n, m int) int {
	return int(math.Pow(float64(n), float64(m)))
}
