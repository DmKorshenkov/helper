package main

import (
	"fmt"
	"strings"
)

// type ymd[K comparable, V any] map[int]map[int]map[int]map[K][]V
// type ymd[K comparable, V any] map[int]map[int]map[int]map[K]V
// type sliceType[S any] []S
// type k[K comparable, V any] map[K]V
func main() {
	helper()
}

func helper() {
	func(month string) int {
		var months = []string{"январь", "февраль", "март", "апрель", "май", "июнь", "июль", "август", "сентябрь", "октябрь", "ноябрь", "декабрь"}
		month = strings.TrimSpace(strings.ToLower(month))
		var a = int(0)
		for in, m := range months {
			if strings.Compare(month, m) == 0 {
				a = in + 1
				break
			}
		}
		fmt.Println(a)
		return a
	}("декабрь")
}
