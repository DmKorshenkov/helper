package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/DmKorshenkov/helper/bot/ev"
	"github.com/DmKorshenkov/helper/bot/sl"
)

func main() {
	os.Chdir("../DataBase")
	data, _ := os.ReadFile("tmp.txt")
	sb := strings.Builder{}

	for _, b := range data {
		sb.WriteByte(b)
	}

	slice := sl.SplitSlashN(sb.String())
	var day = ev.EnergyValue{}
	_ = day
	for _, element := range slice {
		tmp := sl.SplitSpaceTrim(element)
		//	fmt.Println(tmp)

		if eat, ok := sl.GetValInJson[string, ev.EnergyValue](strings.ToLower(tmp[0]), "product.json"); ok {
			//fmt.Println(eat.StringEv())
			if sl.CheckNumber(tmp[1]) {
				eat.SetPortion(sl.ParF(tmp[1]))
				fmt.Println(eat.Str())
			}
		} else {
			fmt.Println("not found")
		}

		tmp = nil
	}
}
