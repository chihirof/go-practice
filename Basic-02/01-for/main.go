package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum2 := 1
	// 初期化と後処理ステートメントの記述は任意。セミコロンも省略可能
	// ループ条件を省略すると、無限ループになる
	for sum2 < 1000 {
		sum2 += sum2
	}
	fmt.Println(sum2)
}
