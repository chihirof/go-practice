package main

import "fmt"

// 定数は const キーワードを使って宣言する
const Pi = 3.14

func main() {
	// 定数の場合は、 := を使って宣言できない
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	// 定数は、文字、文字列、boolean、数値のみで使える
	const Truth = true
	fmt.Println("Go rules?", Truth)

	numericConstants()
}

func numericConstants() {
	const Big = 1 << 100
	const Small = Big >> 99
	const Mid = 100

	// 数値の定数は、高精度な値。型のない定数はその状況によって必要な型をとる
	fmt.Println(Small)
	fmt.Println(needInt(Small))
	fmt.Println(needFload(Small))
	fmt.Println(needFload(Big))
	fmt.Println(needFload(Mid))
}

func needInt(x int) int           { return x*10 + 1 }
func needFload(x float64) float64 { return x * 0.1 }
