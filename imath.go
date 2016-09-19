// Package imath provides some integer math constants and functions.
package imath

// Go does not provide constants for maximum and minimum int and unit values.
// These values can be calulated without knowing their implementation size.
// The section of the go language specification on numeric types
//     https://golang.org/ref/spec#Numeric_types
// contains the necessary information:
//
//     The value of an n-bit integer is n bits wide and represented using two's complement arithmetic.
//
//     There is also a set of predeclared numeric types with implementation-specific sizes:
//
//     uint     either 32 or 64 bits
//     int      same size as uint
const (
	MinUint uint = 0                 // binary: all zeroes
	MaxUint      = ^MinUint          // binary: all ones
	MaxInt       = int(MaxUint >> 1) // binary: all ones except high bit
	MinInt       = ^MaxInt           // binary: all zeroes except high bit
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
