// プログラムはmainパッケージから開始される。
package main

// goのプログラムはパッケージで構成する
import (
	"fmt"
	"math/rand"
)

// チュートリアルの環境で実行するプログラムは常に同じ環境で実行されますので、擬似乱数を返す rand.Intn はいつも同じ数を返します。
func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}
