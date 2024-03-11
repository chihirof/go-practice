package main

import (
	"fmt"
	"time"
)

func main() {
	// selectSample()

	defaultSelect()
}

func selectSample() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		// いずれかのcaseが準備できるまでブロックし、準備ができたcaseを実行する。
		// もし複数のcaseの準備ができている場合、caseはランダムに選択される。
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func defaultSelect() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
			// どのcaseの準備もできていない場合はdefaultが実行される。
		default:
			fmt.Println("     .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
