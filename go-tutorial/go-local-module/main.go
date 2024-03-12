package main

import (
	"fmt"
	"local-module/subpkg1"
	"local-module/subpkg2"
	"log"
)

// https://zenn.dev/nobonobo/articles/4fb018a24f9ee9

func main() {
	message := subpkg1.Subpkg1()
	fmt.Println(message)

	message2 := subpkg2.Subpkg2()
	fmt.Println(message2)

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	msg, err := subpkg1.Hello("")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msg)
}
