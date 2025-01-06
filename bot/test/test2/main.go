package main

import (
	"fmt"
)

// type ymd[K comparable, V any] map[int]map[int]map[int]map[K][]V
type ymd[K comparable, V any] map[int]map[int]map[int]map[K]V
type sliceType[S any] []S
type k[K comparable, V any] map[K]V

func main() {
	fmt.Println("he's alive")
	var mp ymd[int, sliceType[int]]
	var sl sliceType[int]

	_ = mp
	_ = sl
}
