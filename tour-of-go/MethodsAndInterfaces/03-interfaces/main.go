package main

import (
	"fmt"
	"math"
)

func main() {

	var a Abser
	f := MyFloat(-math.Sqrt2)
	// インターフェースの値は、基底型となる具体的な型の値を保持し、メソッドを呼び出すとその基底型の同じ名前のメソッドが実行される
	a = f

	fmt.Println(a.Abs())

	var i I

	// 具体的な値 t がnil のパターン（nilのtを保持するインターフェイス i 自体はnilではない）
	var t *T
	i = t
	i.M()

	i = &T{"hello"}
	i.M()

	// インターフェースの値自体がnilの場合は、ランタイムエラー（呼び出す具体的なメソッドを示す型がインターフェース内に存在しない）
	//var i2 I
	// i2.M()

}

type Abser interface {
	Abs() float64
}

// 型にメソッドを実装していくことによって、インターフェースを実装する（満たす）。明示的に宣言する必要はない(implementsなど)
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type I interface {
	M()
}

type T struct {
	S string
}

// 具体的な値がnilの場合、メソッドはnilをレシーバーとして呼び出される。
// Goではnilをレシーバーとして呼び出されても適切に処理するメソッドを記述するのが一般的。
func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}
