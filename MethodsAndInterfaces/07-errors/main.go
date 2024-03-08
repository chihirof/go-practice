package main

import (
	"fmt"
	"time"
)

func main() {
	// エラーがnilかどうかを判定することでエラーハンドリングする
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

// Goはエラーの状態をerror値で表現する
// error型はfmt.Stringerに似た組み込みのインターフェース
// type error interface {
// 	Error() string
// }

type MyError struct {
	When time.Time
	What string
}

func (e MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return MyError{
		time.Now(),
		"it didn't work",
	}
}
