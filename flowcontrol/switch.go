package main

import (
	"fmt"
	"runtime"
)

// breakは必要ない（自動的に提供される）
func main() {
	fmt.Print("Fo runs on ")

	switch os ;= runtime,GOOS; os {
	case "darwin":
		fmt.PrintIn("OS X.")
	case "linux":
		fmt.PrintIn("Linux. ")
	default:
		fmt.Printf("%s.\n", os)
	}
}