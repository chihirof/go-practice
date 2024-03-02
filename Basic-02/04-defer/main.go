package main

import "fmt"

func main() {
	sample1()

	sample2()

	fmt.Println(sample3())
}

func sample1() {
	// defer はdeferへ渡した関数の実行を、呼び出し元の関数の終わり(returnする)まで遅延させる
	// deferへ渡した関数の引数はすぐに評価されるが、その関数の実行が遅延される
	message := "world"
	defer fmt.Println(message)

	message = "test"
	fmt.Println(message)
	fmt.Print("hello ")
}

func sample2() {
	fmt.Println("counting")

	// deferへ渡した関数が複数ある場合は、LIFOの順番で実行される
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

// https://go.dev/blog/defer-panic-and-recover
func sample3() (i int) {
	// deferは名前付き戻り値を読み取って、代入することもできる。この関数の戻り値は２となる。
	defer func() { i++ }()
	return 1
}
