package main

// factoredインポートステートメント。この書き方の方が良いスタイル
import (
	"fmt"
	"math/rand"
)

func main() {
	// 大文字で始まる名前は外部のパッケージから参照できる名前
	fmt.Println(rand.Intn(10))
}
