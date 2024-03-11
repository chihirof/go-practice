package main

import "fmt"

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	// 受信時に2つ目のパラメータを受け取ることができ、チャネルがクローズされているかどうかがわかる
	v, ok := <-c
	fmt.Println(v, ok)
	// チャネルが閉じられるまでチャネルから値を繰り返し受信し続ける
	for i := range c {
		fmt.Println(i)
	}

}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	// これ以上の送信する値がないことを示すため、チャネルをcloseすることができる（送る側だけがcloseすること）
	close(c)
}
