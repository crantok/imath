# Package imath

Package imath provides some integer math constants

    // Constants stolen from https://groups.google.com/d/msg/golang-nuts/a9PitPAHSSU/ziQw1-QHw3EJ
    const (
    	MaxUint = ^uint(0)
    	MinUint = 0
    	MaxInt  = int(MaxUint >> 1)
    	MinInt  = -MaxInt - 1
    )

and functions.

    func Min(a ...int) int
    func Max(a ...int) int
    func Abs(a int) int
