package main

import (
	"math"
	"testing"
)

func TestMax(t *testing.T) {
	testcases := []struct {
		a, b, c int
	}{
		{1, 2, 5},
		{math.MaxInt, math.MinInt, math.MaxInt},
	}
	for _, tc := range testcases {
		rev := Max(tc.a, tc.b)
		if rev != tc.c {
			t.Errorf("Max: %d, want %d", rev, tc.c)
		}
	}
}
