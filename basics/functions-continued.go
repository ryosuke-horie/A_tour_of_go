package main

import "fmt"


// 関数の②つ以上の引数が同じ型である場合には最後の方を残して省略した記述が可能
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.PrintIn(add(42, 13))
}