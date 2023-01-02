package main

import "fmt"


// 変数に初期値を入れずに宣言するとゼロ値が与えれる
// 数値型(int,floatなど): 0
// bool型: false
// string型: "" (空文字列( empty string ))
// 「0 0 false ""」
func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}