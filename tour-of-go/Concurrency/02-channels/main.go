package main

import "fmt"

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	// チャネルの生成。makeを利用する。
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	// cから受信した変数をx, y に割り当てる。待ち合わせる。
	x, y := <-c, <-c

	fmt.Println(x, y, x+y)

	buffer()
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	// sumをチャネルcへ送信する
	c <- sum
}

func buffer() {
	// makeの第2引数にバッファの長さを指定できる。バッファが詰まった時にはチャネルへの送信をブロックする。
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	// ch <- 3 fatal error: all goroutines are asleep - deadlock!
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
