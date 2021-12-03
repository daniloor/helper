package helper

import "math"

// IfThenElse returns a if the condition is true, otherwise returns b
func IfThenElse(condition bool, a interface{}, b interface{}) interface{} {
	if condition {
		return a
	}
	return b
}

// Round Up returns the nearest integer value of v rounding up.
//
//Special cases are:
//RoundUp(±0) = ±0
//RoundUp(±Inf) = ±Inf
//RoundUp(NaN) = NaN
func RoundUp(x float64) int {
	if x == 0 || math.IsNaN(x) || math.IsInf(x, 0) {
		return int(x)
	}
	xInt := math.Trunc(x)
	if x == xInt {
		return int(x)
	} else {
		return int(x + 1)
	}
}
