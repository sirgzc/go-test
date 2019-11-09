package main

import "fmt"

func main() {
	p := []int{2, 3, 5, 7, 11, 13}

	fmt.Println("p ==", p)

	fmt.Println("p[1:4] ==", p[1:4])

	fmt.Println("p[:3] ==", p[:3])

	fmt.Println("p[4:] ==", p[4:])

	/**
	p == [2 3 5 7 11 13]
	p[1:4] == [3 5 7]
	p[:3] == [2 3 5]
	p[4:] == [11 13]
	*/
}
