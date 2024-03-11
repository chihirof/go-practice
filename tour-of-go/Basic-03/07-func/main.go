package main

import (
	"fmt"
	"math"
)

func main() {
	// 他の変数のように関数を渡すことができる
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	closure()
}

// 関数の引数にとることができる
func compute(fn func(float64, float64) float64) float64 {
	// 戻り値としても使うことができる
	return fn(3, 4)
}

// Goの関数はクロージャ。それ自身の外部から変数を参照する関数値。adder()はクロージャを返す。
func adder() func(int) int {
	sum := 0
	// この関数はsumにアクセスして、sumを変えることができる。sumへバインドされている。
	return func(x int) int {
		sum += x
		return sum
	}
}

func closure() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
}
