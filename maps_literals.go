package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func main() {
	var m = map[string]Vertex{
		"bell": {2, 3},
	}

	fmt.Println(m)

	/**
	map[bell:{2 3}]
	*/
}
