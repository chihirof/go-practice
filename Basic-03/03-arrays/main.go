package main

import "fmt"

func main() {
	// 型intの長さ2の配列を表す。配列の長さは型の一部であり、サイズを変えることはできない。
	var a [2]int
	a[0] = 10
	a[1] = 20
	fmt.Println((a))

	var b = [3]int{2, 3, 5}
	fmt.Println(b)

}
