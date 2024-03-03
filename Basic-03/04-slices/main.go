package main

import "fmt"

func main() {
	// 配列は固定長
	a := [6]int{1, 2, 3, 4, 5, 6}

	// 型 []T は 型Tのスライスを表す。スライスは可変長。単に元の配列の部分列を指し示す。
	var s []int = a[1:4]
	fmt.Println(s)

	var t []int = a[1:5]
	fmt.Println(s)

	// スライスを変更すると、元の配列も、他のスライスも更新される
	s[1] = 200
	fmt.Println("a(array): ", a)
	fmt.Println("s(slice): ", s)
	fmt.Println("t(slice): ", t)

	a[4] = 500
	fmt.Println("a(array): ", a)
	fmt.Println("s(slice): ", s)
	fmt.Println("t(slice): ", t)
}
