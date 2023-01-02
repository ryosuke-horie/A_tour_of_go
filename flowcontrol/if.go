package main

import (
	"fmt"
	"math"
)

// Tips: 条件式内で変数を書くことも可能（ifスコープでしか使えない）
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}

	return fmt.Sprint(math.Sqrt(x))
}

func main () {
	fmt.PrintIn(sqrt(2), sqrt(-4))
}