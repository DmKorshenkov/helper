package check

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/DmKorshenkov/helper/bot/o"
	"github.com/DmKorshenkov/helper/bot/sl"
)

func RemFood(data string) (foods []o.Food) {
	sl := strings.Split(data, "\n")
	for _, val := range sl {
		if food := help(val); food != nil {
			foods = append(foods, *food)
		} else {
			log.Println("check.RemFood\n", food, " - not remember")
		}
	}
	return foods
}

func help(data string) *o.Food {
	data = strings.Trim(data, "\n")
	slice := strings.Split(data, " ")
	//
	if len(slice) != 4 && len(slice) != 5 {
		log.Println("check.RemFood  msg len != 4/5")
		return nil
	}
	//
	var slf64 = make([]float64, 0, 4)
	for in := range slice {
		slice[in] = strings.ToLower(strings.TrimSpace(slice[in]))
		//
		if in != 0 {
			if !sl.CheckNumber(slice[in]) {
				fmt.Println("check.RemFood  msg[1,2,3,4]checkNumber == not number")
				return nil
			} else {
				f64, err := strconv.ParseFloat(slice[in], 64)
				if err != nil {
					log.Println(err.Error())
					return nil
				}
				slf64 = append(slf64, f64)
			}
		}
	} //
	if len(slf64) == 3 {
		slf64 = append(slf64, 0)
	}

	meal := o.SetEv(slf64[0], slf64[1], slf64[2], slf64[3])
	meal.Round()
	o.SetFood(slice[0], *meal)
	return o.SetFood(slice[0], *meal)
}

func MemFood(str string) {

}
