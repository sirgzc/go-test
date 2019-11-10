package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	pow := make([]int, 10)

	for i := range pow {
		pow[i] = 1 << uint(i)
	}

	for j, value := range pow {
		fmt.Printf("pow[%d]==%d\n", j, value)
	}

	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
	/**
	pow[0]==1
	pow[1]==2
	pow[2]==4
	pow[3]==8
	pow[4]==16
	pow[5]==32
	pow[6]==64
	pow[7]==128
	pow[8]==256
	pow[9]==512
	1
	2
	4
	8
	16
	32
	64
	128
	256
	512
	*/
}
