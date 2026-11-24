package main

import "testing"

func TestNthRoot(t *testing.T) {
	cases := []struct {
		n, m, want int
	}{
		{3, 27, 3},
		{4, 69, -1},
		{2, 16, 4},
		{1, 14, 14},
		{5, 1, 1},
		{2, 0, 0},
		{10, 1024, 2},
		{2, 2147483647, -1},
		{3, 1000000000, 1000},
		{0, 5, -1},
	}
	for _, c := range cases {
		if got := NthRoot(c.n, c.m); got != c.want {
			t.Errorf("NthRoot(%d,%d) = %d, want %d", c.n, c.m, got, c.want)
		}
	}
}
