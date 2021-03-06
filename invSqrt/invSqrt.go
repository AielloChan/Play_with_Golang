package main

import (
	"fmt"
	"unsafe"
)

func InvSqrt(x float32) float32 {

	var xhalf = 0.5 * x
	i := *(*int)(unsafe.Pointer(&x))
	i = 0x5f375a86 - (i >> 1)
	x = *(*float32)(unsafe.Pointer(&i))
	x = x * (1.5 - xhalf*x*x)
	x = x * (1.5 - xhalf*x*x)
	x = x * (1.5 - xhalf*x*x)

	return 1 / x
}

func main() {
	fmt.Println(InvSqrt(65535))
}
