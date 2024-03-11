package main

import (
	"fmt"
	"math"
)

// 型にメソットを定義できる
type Vertex struct {
	X, Y float64
}

// 特別なレシーバをとる（v Vertex）。Absメソッドはvという名前のVertex型のレシーバを持つことを意味する
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

// typeにもメソッドを宣言できる
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
