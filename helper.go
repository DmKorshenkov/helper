package helper

import (
	"strings"
	"fmt"
)

func SplitTrimCheck(get string, n int) []string {
	slice := strings.Split(get, " ")
	if len(slice) == n || n == 0 {
		for in, str := range slice {
			slice[in] = strings.TrimSpace(str)
		}
		return slice
	}
	return nil
}

func CheckErr(err error){
	if err != nil {
		log.Println(err.Errors())
	}
}
