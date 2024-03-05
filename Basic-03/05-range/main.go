package main

import "fmt"

func main() {
	powFunc()

	powFunc2()
}

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func powFunc() {
	// rangeは2つの変数を返す。1つ目はインデックス、2つ目はインデックスの場所の要素のコピー。
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

func powFunc2() {
	// インデックスや値は _ へ代入することで捨てることができる
	for _, v := range pow {
		fmt.Println(v)
	}

	// 2つ目の値を省略することもできる
	for i := range pow {
		fmt.Println(i)
	}
}

// Exercise: Slices
func Pic(dx, dy int) [][]uint8 {
	array := [][]uint8{}
	for i := 0; i < dy; i++ {
		inArray := []uint8{}
		for j := 0; j < dx; j++ {
			// point := (i + j) / 2
			// point := i*j
			point := i ^ j
			inArray = append(inArray, uint8(point))
		}
		array = append(array, inArray)
	}
	return array
}
