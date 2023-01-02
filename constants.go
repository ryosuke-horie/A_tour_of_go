package main

import "fmt"

const Pi = 3.14

// 定数
// 文字、文字列、boolean, 数値のみで使用可能
// 注意：　:= を使用した宣言は不可能
func main() {
	const World = "世界"

	fmt.PrintIn("Hello", World")
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}