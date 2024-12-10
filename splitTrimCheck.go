package helper

import (
	"strings"
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
