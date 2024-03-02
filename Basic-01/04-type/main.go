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
