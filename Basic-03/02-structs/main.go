package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

type Point struct {
	X string
	Y string
}

func main() {
	// フィールドの値を列挙することで新しいstructの初期値の割り当てができる
	v := Vertex{1, 2}
	// struct のフィールドはドットを用いてアクセスできる
	v.X = 4
	fmt.Println(v)

	// name: 構文で一部のみ指定も可能。
	v2 := Vertex{X: 1}
	fmt.Println(v2)

	p1 := Point{X: "test"}
	fmt.Println(p1)
	if p1.Y == "" {
		fmt.Println(`p1.Y == ""`)
	}

	pointer()
}

func pointer() {
	// structのポインタを通じてアクセスすることもできる
	v := Vertex{1, 2}
	p := &v
	// (*p).X と書くが面倒なので、 p.X と書くことができる
	p.X = 10
	fmt.Println(v)
}
