package main 

import {
	"fmt"
	"math"
}

// Goでは最初の文字が大文字で始まる名前は外部公開された名前。Piはmathパッケージでエクスポートされている。
func main() {
	// ↓エラーが発生する
	// fmt.Println(math.pi)

	// ↓エラーは発生しない
	fmt.Println(math.Pi)
}