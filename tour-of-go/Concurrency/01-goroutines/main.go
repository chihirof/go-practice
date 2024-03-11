package main

import (
	"fmt"
	"time"
)

// goroutineはGoのランタイムに管理される軽量なスレッド。
func main() {
	// go f(x, y, z) とすることで、新しいgoroutineが実行される。
	// f, x, y, z は実行元のgoroutineで評価され、fの実行は新しいgoroutineで実行される。
	go say("world")
	say("hello")
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
