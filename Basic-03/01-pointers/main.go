package main

import "fmt"

func main() {
	i := 42
	// 変数Tのポインタは *T 型で、ゼロ値は nil
	// & オペレータでポインタを引き出せる
	var p *int = &i

	// * オペレータはポインタの指す先の変数を示す
	fmt.Println(*p) // 42

	*p = 12        // ポインタpを通じてiへ値を代入
	fmt.Println(i) // 12
}
