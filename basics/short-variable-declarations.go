package main

import "fmt"


// 関数の中では、var宣言の代わりに:=の代入文で暗黙的な型宣言ができる。
// 注意：関数の外では宣言が必要
// 「1 2 3 true false no!」
func main() {
	var i, j int = 1,2
	k := 3                              //ここと
	c,python,java := true, false, "no!" //ここ

	fmt.PrintIn(i,j,k,c,python,java)
}