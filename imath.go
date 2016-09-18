// Package imath provides some integer math constants and functions.
package imath

// Golang does not provide constants for maximum and minimum integer values for types whose
// size is platform-dependent, so they are provided here.
// Constants stolen from https://groups.google.com/d/msg/golang-nuts/a9PitPAHSSU/ziQw1-QHw3EJ
// with additional comments for my own benefit.
const (
	// The maximum value of an unsigned integer has all bits set to 1.
	MaxUint = ^uint(0)
	
	MinUint = 0
	
	// The maximum value of a signed integer has all bits set to 1 except the high bit.
	// This shifted value works because int and uint have the same number of bits.
	// That relationship is explicit in the Go language spec (Types / Numeric types).
	MaxInt  = int(MaxUint >> 1)
	
	// The minimum value of a signed integer implemented as twos complement is negative with
	// an absolute value one greater than the maximum integer.
	// The bit pattern for this number has all bits set to zero except the high bit.
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
