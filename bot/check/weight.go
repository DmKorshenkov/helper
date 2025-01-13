package check

import (
	"log"
	"strconv"
	"strings"

	"github.com/DmKorshenkov/helper/bot/o"
	"github.com/DmKorshenkov/helper/bot/sl"
)

func RemWeight(data string) *o.Weight {

	slice := strings.Split(data, " ")
	for in := range slice {
		slice[in] = strings.ToLower(strings.TrimSpace(slice[in]))
		if in == 0 && !sl.CheckNumber(slice[in]) {
			return nil
		}
	}
	if len(slice) == 2 {
		if sl.CheckNumber(slice[0]) {
			w, err := strconv.ParseFloat(slice[0], 64)
			if err != nil {
				log.Println(err.Error())
				return nil
			}
			info := slice[1]
			return o.SetW(w, info)
		} else {
			//CheckNumber== false; return nil
			return nil
		}
	}
	if len(slice) == 1 {
		if sl.CheckNumber(slice[0]) {
			w, err := strconv.ParseFloat(slice[0], 64)
			if err != nil {
				log.Println(err.Error())
				return nil
			}
			return o.SetW(w, "")
		} else {
			//CheckNumber== false; return nil
			return nil
		}
	}
	return nil
}
func CheckMemWeight(str string) {

}
