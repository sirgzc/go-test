package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	time.Sleep(500 * time.Millisecond)
	say("hello")
	go say("aaa")
	time.Sleep(1000 * time.Millisecond)
	say("bbb")
}
