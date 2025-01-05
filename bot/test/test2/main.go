package main

import "fmt"

func main() {
	var sl = make([]int, 0, 10)
	for i := 0; i < cap(sl); i++ {
		sl = append(sl, i)
	}
	fmt.Println(sl[2:4])
}
