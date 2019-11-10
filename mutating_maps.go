package main

import "fmt"

func main(){
	m := make(map[string]int)
	
	m["anwser"] = 42
	fmt.Println(m["anwser"]) //42
	
	m["anwser"] = 48
	fmt.Println(m["anwser"]) // 48
	
	delete(m, "anwser")
	fmt.Println(m["anwser"]) //0
	
	v, ok := m["anwser"]
	fmt.Println("The value:", v, "Present?", ok) //The value: 0 Present? false
}
