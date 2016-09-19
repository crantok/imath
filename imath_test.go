package imath_test

import (
	"math"
	"testing"

	"github.com/crantok/imath"
)

func TestConsts(t *testing.T) {
	if imath.MinUint != 0 {
		t.Errorf("MinUint, expected 0, got %v", imath.MinUint)
	}
	if uint64(imath.MaxUint) != math.MaxUint32 && uint64(imath.MaxUint) != math.MaxUint64 {
		t.Errorf("MaxUint, expected %v or %v, got %v", uint32(math.MaxUint32), uint64(math.MaxUint64), imath.MinUint)
	}
	if int64(imath.MaxInt) != math.MaxInt32 && int64(imath.MaxInt) != math.MaxInt64 {
		t.Errorf("MaxInt, expected %v or %v, got %v", int32(math.MaxInt32), int64(math.MaxInt64), imath.MaxInt)
	}
	if int64(imath.MinInt) != math.MinInt32 && int64(imath.MinInt) != math.MinInt64 {
		t.Errorf("MaxInt, expected %v or %v, got %v", int32(math.MinInt32), int64(math.MinInt64), imath.MinUint)
	}
	if imath.MinInt+imath.MaxInt != -1 {
		t.Errorf("MinInt + MaxInt, expected -1, got %v", imath.MinInt+imath.MaxInt)
	}
}

func TestFuncs(t *testing.T) {
	min := -32
	absMin := 32
	max := 35
	vals := []int{3, min, max, 0}

	x := imath.Min(vals...)
	if x != min {
		t.Errorf("Min(%v), expected %v, got %v", vals, min, x)
	}

	x = imath.Max(vals...)
	if x != max {
		t.Errorf("Max(%v), expected %v, got %v", vals, max, x)
	}

	x = imath.Abs(min)
	if x != absMin {
		t.Errorf("Abs(%v), expected %v, got %v", min, absMin, x)
	}

}
