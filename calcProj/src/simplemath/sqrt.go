package simplemath

import (
	"math"
)

// Sqrt 开方运算
func Sqrt(i int) int {
    v := math.Sqrt(float64(i))
    return int(v)
}