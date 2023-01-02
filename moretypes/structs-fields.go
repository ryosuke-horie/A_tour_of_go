package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

// structのフィールドはドットでアクセスする
func main() {
	v := Vertex{1,2}
	v.X = 4
	fmt.PrintIn(v.X)
}