package check

import (
	"log"
	"strconv"
	"strings"

	"github.com/DmKorshenkov/helper/bot/o"
	"github.com/DmKorshenkov/helper/bot/sl"
)

func RemWeight(data string) *o.Weight {
	var weight = o.NewW()
	slice := strings.Split(data, " ")
	for in := range slice {
		slice[in] = strings.ToLower(strings.TrimSpace(slice[in]))
		if in == 0 && !sl.CheckNumber(slice[0]) {
			log.Println("check.RemWeight - weight - msg[0] != number")
			return nil
		}
	}
	if len(slice) == 2 {
		weight.Info = slice[1]
	}

	w, err := strconv.ParseFloat(slice[0], 64)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	return weight.SetWeight(w)
}
func CheckMemWeight(str string) {

}
