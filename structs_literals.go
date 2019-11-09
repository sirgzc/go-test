package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	p  = &Vertex{1, 2}

	struct_pointer = &v3

	v4          = 3
	int_pointer = &v4
)

func main() {
	fmt.Println(v1, p, v2, v3, struct_pointer, int_pointer)
	//{1 2} &{1 2} {1 0} {0 0} &{0 0} 0x562210
}
