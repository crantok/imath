// Package imath provides some integer math constants and functions.
package imath

// Constants stolen from https://groups.google.com/d/msg/golang-nuts/a9PitPAHSSU/ziQw1-QHw3EJ
const (
	MaxUint = ^uint(0)
	MinUint = 0
	MaxInt  = int(MaxUint >> 1)
	MinInt  = -MaxInt - 1
)

// Min gets the minimum value from a set of ints. When called with zero
// arguments, Min returns MaxInt.
func Min(a ...int) int {

	min := MaxInt

	for _, v := range a {
		if v < min {
			min = v
		}
	}

	return min
}

// Max gets the maximum value from a set of ints. When called with zero
// arguments, Max returns MinInt.
func Max(a ...int) int {

	max := MinInt

	for _, v := range a {
		if v > max {
			max = v
		}
	}

	return max
}

// Abs returns the absolute value of an int.
func Abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}
