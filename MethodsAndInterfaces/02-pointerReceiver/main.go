package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// ポインタレシーバ。変数レシーバは元のVertex変数のコピーを宣言する。元のVertex変数を更新したい場合はポインタレシーバにしないといけない
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Vertex) Scale2(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// ポインタレシーバを使う理由
// 1. メソッドがレシーバが指す先の変数を変更するため
// 2. メソッドの呼び出しごとに変数のコピーを避けるため

func main() {
	v := Vertex{3, 4}
	// Scaleのレシーバはポインタレシーバだが、変数、もしくは、ポインタのどちらでもOK
	// Goが (&v2).Scale(10) と解釈する
	v.Scale(10)
	fmt.Println(v.Abs()) // 50

	v2 := Vertex{3, 4}
	v2.Scale2(10)
	fmt.Println(v2.Abs()) // 10

	p := &Vertex{3, 4}
	// Scale2のレシーバとして、ポインタをとることもできる。
	// Go が (*p).Scale2(10) と解釈する
	p.Scale2(10)
	fmt.Println(p.Abs())
}
