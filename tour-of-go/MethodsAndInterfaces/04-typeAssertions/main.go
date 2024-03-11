package main

import "fmt"

func main() {
	var i interface{} = "hello"

	// 型アサーション。インターフェースの値の基になる具体的な値を利用する手段を提供する。
	// インターフェースiが具体的な型stringを保持し、基になるstringの値を変数sに代入する。もしiがstringを保持していない場合、panicを引き起こす
	s := i.(string)
	fmt.Println(s)

	// 特定の型を保有しているかどうかは2個目の戻り値でわかる
	s, ok := i.(string)
	fmt.Println(s, ok)

	// okはfalseとなり、fはfloat64のゼロ値になり、panicは起きない
	f, ok := i.(float64)
	fmt.Println(f, ok)

}
