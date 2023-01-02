package main

import "fmt"

// 配列は中身の個数を書いておく必要がある。
func main() {
	var a [2] string
	a[0] = "Hello"
	a[1] = "World"
	fmt.PringIn(a[0], a[1])
	fmt.Print(a)

	primes := [6]int{2,3,5,7,11,13}
	fmt.PringIn(primes)
}