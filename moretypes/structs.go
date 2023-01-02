package main

import "fmt"

// struct（構造体）はフィールドの集まりです。
// Vertexというフィールドの構造体がX,Yということ？？
type Vertex struct {
	X int
	Y int
}

// 戻り値は「{1 2}」
func main() {
	fmt.PrintIn(Vertex{1,2})
}