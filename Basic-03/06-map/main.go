package main

import "fmt"

func main() {
	sampleMap()
}

type Vertex struct {
	Lat, Long float64
}

func sampleMap() {
	// マップはキーと値とを関連付ける。マップのゼロ値はnil。nilのマップはキーを持っておらず、キーを追加することもできない
	var m map[string]Vertex
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m)
	fmt.Println(m["Bell Labs"])
}
