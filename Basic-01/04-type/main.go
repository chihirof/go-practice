package main

import (
	"fmt"
	"math/cmplx"
)

// import文と同じようにまとめて宣言可能
var (
	ToBe              = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12)
	c      rune       = 'c' // rune は int32 の別名。文字そのものを表し、Unicodeのコードポイントを表す
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	fmt.Printf("Type: %T Value: %v\n", c, c)

	zeroValue()

	conversionType()

	typeInterface()
}

func zeroValue() {
	// 数値型のゼロ値は0、bool型はfalse、string型は "" になる
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

func conversionType() {
	// 変数v、型Tがあった時、T(v)は変数vをT型へ変換する。Goでは明示的な変換が必要
	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)

	fmt.Println(i, f, u)
}

func typeInterface() {
	// 明治的な型を指定せずに変数を宣言すると、変数の型は右辺の変数から型推論される。
	var i int
	j := i
	fmt.Printf("%T\n", i)
	fmt.Printf("%T\n", j)

	// 右辺が数値である場合、右辺の定数の精度に基づいて、int、float64、complex128 の型になる。
	k := 42
	l := 1.5
	m := 0.867 + 0.5i
	fmt.Printf("%T\n", k)
	fmt.Printf("%T\n", l)
	fmt.Printf("%T\n", m)

	n := "hello"
	o := 'a'
	fmt.Printf("%T\n", n)
	fmt.Printf("%T\n", o)
}
