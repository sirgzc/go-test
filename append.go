package main

import "fmt"

func main() {
	var a []int
	printSlice("a", a)

	a = append(a, 0)
	printSlice("a", a)

	a = append(a, 1)
	printSlice("a", a)

	a = append(a, 2, 3, 4)
	printSlice("a", a)

	/**
	a len=0 cap=0 []
	a len=1 cap=1 [0]
	a len=2 cap=2 [0 1]
	a len=5 cap=6 [0 1 2 3 4]
	*/
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
