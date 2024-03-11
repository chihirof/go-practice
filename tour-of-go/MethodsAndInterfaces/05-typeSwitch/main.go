package main

import "fmt"

func main() {
	do(32)
	do("Hello")
	do(true)
	do(0.2)
}

func do(i interface{}) {
	// 型Switchを使うことで、いくつかの型アサーションを直列に使用できる
	// case は型を指定し、それらの値は指定されたインターフェースの値が保持する値の型と比較される
	switch v := i.(type) { // i.(T) と同じ構文だが、特定の型Tは type に置き換えられる
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}
