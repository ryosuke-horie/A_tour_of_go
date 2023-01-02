package main

// ここの書き方について。もちろん並べてそれぞれimportもできるけど↓の書き方のほうがより良い。
import (
	"fmt"
	"math"
)


// 実行結果は「Now you have 2.6457513110645907 problems.」
func main() {
	fmt.Printf("Now you have %g problems. \n", math.Sqrt(7))
}