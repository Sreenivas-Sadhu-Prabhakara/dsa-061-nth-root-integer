package main

import "fmt"

// NthRoot returns the integer n-th root of m when m is a perfect n-th power,
// and -1 otherwise. Binary search over candidate roots, with a power routine
// that caps out instead of overflowing.
func NthRoot(n, m int) int {
	if n <= 0 || m < 0 {
		return -1
	}
	if m == 0 || m == 1 {
		return m
	}
	lo, hi := 1, m
	for lo <= hi {
		mid := lo + (hi-lo)/2
		switch p := cappedPower(mid, n, m); {
		case p == m:
			return mid
		case p < m:
			lo = mid + 1
		default:
			hi = mid - 1
		}
	}
	return -1
}

// cappedPower returns x^n, or cap+1 as soon as the product exceeds cap —
// which both avoids overflow and is all the binary search needs to know.
func cappedPower(x, n, cap int) int {
	res := 1
	for i := 0; i < n; i++ {
		if res > cap/x {
			return cap + 1
		}
		res *= x
	}
	return res
}

func main() {
	fmt.Println(NthRoot(3, 27), NthRoot(4, 69), NthRoot(2, 1_000_000))
}
