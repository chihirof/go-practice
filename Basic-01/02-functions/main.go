package main

import "fmt"

func main() {
	fmt.Println(add(42, 13))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(21))
}

func add(x int, y int) int {
	return x + y
}

// 2つ以上の引数が同じ型なら、最初の引数の型を省略可能
func add2(x, y int) int {
	return x + y
}

// 複数の戻り値を返すことができる
func swap(x, y string) (string, string) {
	return y, x
}

// 戻り値に名前をつけることができる。関数の最初で定義した変数名として扱われる。
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	// 名前をつけた戻り値の変数を使うと、何も書かなくても戻り値をreturnできる（naked return）。
	// 読みやすさの観点から短い関数でのみ利用すべき
	return
}
