package absmax

import "math"

// AbsInt64 returns the absolute value of n.
//
// Branchless: see https://books.google.com.au/books?id=
// VicPJYM0I5QC&lpg=PA18&ots=2o-SROAuXq&dq=hackers%20delight%20absolute&pg=PA18
// #v=onepage&q=hackers%20delight%20absolute&f=false
func AbsInt64(n int64) int64 {
	y := n >> 63
	return (n ^ y) - y
}

// AbsInt returns the absolute value of n.
func AbsInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// AbsInt32 returns the absolute value of n.
func AbsInt32(n int32) int32 {
	y := n >> 31
	return (n ^ y) - y
}

// AbsInt16 returns the absolute value of n.
func AbsInt16(n int16) int16 {
	y := n >> 15
	return (n ^ y) - y
}

// AbsInt8 returns the absolute value of n.
func AbsInt8(n int8) int8 {
	y := n >> 7
	return (n ^ y) - y
}

// AbsFloat64 returns the absolute value of n.
func AbsFloat64(n float64) float64 {
	return math.Abs(n)
}

// AbsFloat32 returns the absolute value of n.
func AbsFloat32(n float32) float32 {
	if n < 0 {
		return -n
	}
	return n
}
