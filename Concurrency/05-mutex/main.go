package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}

type SafeCounter struct {
	mu sync.Mutex // 排他制御
	v  map[string]int
}

func (c *SafeCounter) Inc(key string) {
	// LockとUnlockで囲むことで、排他制御で実行するコードを定義できる
	c.mu.Lock()
	c.v[key]++
	c.mu.Unlock()
}

func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// deferを使うことで、mutexがUnlockされることを保証する。
	defer c.mu.Unlock()
	return c.v[key]
}
