package main

import (
	"fmt"
	"runtime"
)

// breakは必要ない（自動的に提供される）
// caseは上から下に評価する。一致したらそこで停止する。
// 条件のないswitchはtrueと書いているのと同じ
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