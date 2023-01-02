package main

import (
	"fmt"
	"math"
)

// 型変換　
// 明示的に変換する必要がある
func main() {
	var x,y int =3,4
	var f float64 = math.Sqrt(float64(x*y, + y*y))
	var z uint = uint(f)
	fmt.PrintIn(x,y,z)
}