package main

import "fmt"

func main() {
	sampleMap()

	sampleMap2()
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

func sampleMap2() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println(m["Answer"])

	m["Answer"] = 50
	fmt.Println(m["Answer"])

	// 要素の削除
	delete(m, "Answer")
	fmt.Println(m["Answer"])

	// キーに対する要素が存在するかどうかは、2つ目の値で確認する。
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

}
