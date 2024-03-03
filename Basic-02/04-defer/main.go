package main

import "fmt"

func main() {
	sample1()

	sample2()

	fmt.Println(sample3())

	f()
	fmt.Println("Returned normally from f.")
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

func f() {
	defer func() {
		// パニックに陥ったゴルーチンの制御を回復する組み込み関数。遅延関数の中でのみ有効。
		// 通常の実行中は、nilを返す。現在のゴルーチンがパニックに陥っている時はpanicに与えられた値を読み込み、通常の処理を開始する
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		// 通常の制御の流れを止めて、パニックを開始する組み込み関数。panicが呼び出されると、g()の実行は停止し、遅延関数を実行して呼び出し元に戻る
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}
