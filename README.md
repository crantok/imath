# Package imath

Package imath provides some integer math constants

    const (
    	MinUint uint = 0                 // binary: all zeroes
    	MaxUint      = ^MinUint          // binary: all ones
    	MaxInt       = int(MaxUint >> 1) // binary: all ones except high bit
    	MinInt       = ^MaxInt           // binary: all zeroes except high bit
    )

and functions.

    func Min(a ...int) int
    func Max(a ...int) int
    func Abs(a int) int
