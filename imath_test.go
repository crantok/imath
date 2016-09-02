package imath_test

import (
	"testing"

	"github.com/crantok/imath"
)

func Test(t *testing.T) {
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
