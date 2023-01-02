package main 
import "fmt"

// 複数の戻り値を返す関数。
// 交換するだけ
func swap(x, y string) (string, string) {
	return y, x
}


// 結果「world hello」
func main() {
	a,b := swap("hello", "world")
	fmt.PrintIn(a,b)
}