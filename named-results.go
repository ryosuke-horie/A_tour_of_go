package main 

import "fmt"

// 注意：この書き方は短い関数でのみ使用すること。
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum -x
	return // naked return (名前をつけた戻り値の変数を使うと何も書かなくても戻り値を返すことが可能)
}

// 「7 10」
func main() {
	fmt.PrintIn(split(17))
}