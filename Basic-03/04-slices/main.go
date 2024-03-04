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

	// これは [3]bool{true, false, false} と同様の配列を作り、それを参照するスライスを生成している
	d := []bool{true, false, false}
	fmt.Println(d)

	l := []int{2, 3, 5, 7, 11, 13}
	// スライスの範囲は下限が0、上限がスライスの長さ。以下のスライス式は全て同じ
	fmt.Println(l[0:6])
	fmt.Println(l[:6])
	fmt.Println(l[0:])
	fmt.Println(l[:])

	sliceLengthAndCapacity()

	sliceNil()

	useMake()

	sliceOfSlice()

	appendToSlice()

	blog()
}

func sliceLengthAndCapacity() {
	s := []int{2, 3, 5, 7, 11, 13}
	// スライスは長さと容量を持つ
	// スライスの長さは要素の数
	// スライスの容量は、スライスの最初の要素から考えて、元となる配列の要素の数
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	s = s[:0]
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	s = s[:4]
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	s = s[2:]
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func sliceNil() {
	// スライスのゼロ値はnil
	var s []int
	fmt.Println(s, len(s), cap(s))
	fmt.Printf("%v\n", s)
	if s == nil {
		fmt.Println("nil")
	}
}

func useMake() {
	// makeはゼロ値化された配列を割り当て、その配列を指すスライスを返す
	a := make([]int, 5)
	fmt.Printf("len=%d cap=%d %v\n", len(a), cap(a), a)

	// makeの第3引数にスライスの容量を指定できる(スライス: [], 配列:[0, 0, 0, 0, 0])
	b := make([]int, 0, 5)
	fmt.Printf("len=%d cap=%d %v\n", len(b), cap(b), b)

	// スライス [0, 0], 配列: [0, 0, 0, 0, 0]
	c := b[:2]
	fmt.Printf("len=%d cap=%d %v\n", len(c), cap(c), c)

	// スライス [0, 0, 0], 配列: [0, 0, 0]
	d := c[2:5]
	fmt.Printf("len=%d cap=%d %v\n", len(d), cap(d), d)
}

func sliceOfSlice() {
	// スライスは他のスライスを含む任意の方を含むことができる
	board := [][]string{
		[]string{"-", "-", "-"},
		[]string{"-", "-", "-"},
		[]string{"-", "-", "-"},
	}
	fmt.Println(board)
}

func appendToSlice() {
	var s []int
	printSlice(s)

	// Goの組み込み関数の append を使うことで新しい要素を足せる
	s = append(s, 10)
	printSlice(s)

	s = append(s, 20, 30, 40)
	printSlice(s)

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// https://go.dev/blog/slices-intro
func blog() {
	s := make([]byte, 5, 5) // []byte{0, 0, 0, 0, 0}
	fmt.Println(len(s), cap(s))

	s = s[2:4]                  // []byte{0, 0}
	fmt.Println(len(s), cap(s)) // len:2, cap:3

	s = s[:cap(s)] // []byte(0, 0, 0)
	fmt.Println(len(s), cap(s))
}
