package main

import (
	"fmt"
	"unicode"
)

func main() {
	str := "12.ф2"
	fmt.Println(helper(str))

}

func helper(slice ...string) bool {
	for _, str := range slice {
		for _, r := range str {
			if !unicode.IsDigit(r) && r != '.' {
				return false
			}
		}
	}
	return true
}
