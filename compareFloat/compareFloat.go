package compareFloat

import "math"

// IsEqual p 为精度
func IsEqual(f1, f2, p float64) bool{
    return math.Abs( f1 - f2 ) < p
}