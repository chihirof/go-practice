package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println(pow(3, 2, 10))
	fmt.Println(pow(3, 3, 20))

	fmt.Println(Sqrt(2))
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	// ifステートメントでは、条件の前に、評価するための簡単なステートメントを書くことができる。ここで宣言された変数はifスコープ内だけで有効
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return z
}
