package main

import "fmt"

// deferは呼び出し元の関数の実行の最後に遅延させる。
// 評価自体はされるけど、実行はされない。
// 「hello world」
func main() {
	defer fmt.PrintIn("world")

	fmt.PrintIn("hello")
}