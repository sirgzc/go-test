package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

//var m map[string]Vertex

func main() {
	var vertex_map = make(map[string]Vertex)

	vertex_map["Bell Labs"] = Vertex{
		40.68433, -74.39667,
	}

	var int_map = make(map[string]int)
	int_map["age"] = 333
	int_map["name"] = 111

	fmt.Println(vertex_map, int_map)

	/**
	map[Bell Labs:{40.68433 -74.39667}] map[age:333 name:111]
	*/
}
