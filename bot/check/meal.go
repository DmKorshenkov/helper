package check

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/DmKorshenkov/helper/bot/opr"
	"github.com/DmKorshenkov/helper/bot/sl"
)

func CheckRemMeal(data string) (meal []opr.Prod) {
	slice := strings.Split(data, "\n")
	fmt.Println("!!!!!!!", slice, len(slice))

	for in := range slice {
		slice[in] = strings.ToLower(strings.Trim(slice[in], "\n"))
		s := strings.Split(slice[in], string(rune(' ')))
		
		if len(s) != 2 {
			fmt.Println("CheckRemML meal len == false")
			return nil
		}
		s[0] = strings.TrimSpace(s[0])
		s[1] = strings.TrimSpace(s[1])
		if opr.RemProd(s[0]) == nil {
			fmt.Println("CheckRemML RemProd(s[0]) == nil ")
		}
		if !sl.CheckNumber(s[1]) {
			fmt.Println("CheckRemML CheckNumber(s[1]) == false")
			return nil
		} else {
			f64, _ := strconv.ParseFloat(s[1], 64)
			meal = append(meal, opr.Prod{Name: s[0], Weight: f64})
			//if err != nil
		}
	}
	return meal
}

func CheckMemMl(str string) {

}
