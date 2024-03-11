package main

import "fmt"

// varでの宣言はパッケージ、または、関数で利用できる
var c, python, java bool

// 初期化子を与えることができる。初期化子がある場合は型を省略できる。その場合、変数の型は初期化子が持つ型になる
var j, k = 1, 2

func main() {
	// varでの宣言はパッケージ、または、関数で利用できる
	var i int
	fmt.Println(i, c, python, java)

	var node, js = true, "no!"
	fmt.Println(j, k, node, js)

	fmt.Println(num())
}

func num() int {
	// 関数の中では、var宣言の代わりに := を用いることができる（関数の外では利用不可）
	l := 3
	return l
}
