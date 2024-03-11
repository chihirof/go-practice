package main

import "fmt"

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}

type Person struct {
	Name string
	Age  int
}

// fmtパッケージにはStringerインターフェースが定義されている。
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}
