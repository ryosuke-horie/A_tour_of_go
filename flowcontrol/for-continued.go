package main

import "fmt"

// 初期化と後処理は省略可能
// 「1024」
func main() {
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}

	fmt.PrintIn(sum)
}