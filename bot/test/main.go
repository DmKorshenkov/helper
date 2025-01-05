package main

import (
	"encoding/json"
	"fmt"
	"os"

	ev "github.com/DmKorshenkov/helper/bot/ev"
	opr "github.com/DmKorshenkov/helper/bot/opr"
	ymd "github.com/DmKorshenkov/helper/bot/ymd"
)

func main() {
	os.Chdir(".././o")

	products := opr.RemAllProd()
	fmt.Println(products)

}

func mealTake(o2 []opr.Object) {
	o := o2
	for in, food := range o {
		//	fmt.Println(food)
		if check := meal(food); check != nil {
			o[in] = *check
		} else {
			fmt.Println("check==nil")
		}

	}
	//все продукты в o
	//fmt.Println(o)
	sumev := ev.EnergyValue{}
	for _, food := range o {
		sumev.SumEv(food.EnergyValue)
	}
	o = append(o, *opr.SetO("БЖУ", sumev))
	//MemMeal(o)
	opr.RemRate()
	//MemRate(RemRate().DiffEv(o[len(o)-1].EnergyValue))
}

func meal(food opr.Object) *opr.Object {
	if f := opr.RemProd(food.Name); f != nil {
		f.SetOneGram().SetPortion(food.EnergyValue.W.Weight)
		return (opr.SetO(food.Name, *f))
	} else {
		fmt.Println("f==nil")
		return nil
	}

}

func MemMeal(e []opr.Object) {

	var mp = make(map[int]map[int]map[int]map[int][]opr.Object)
	key := 1
	data, _ := os.ReadFile("mealtake.json")
	if len(data) != 0 {
		json.Unmarshal(data, &mp)

		if y, m, d := ymd.SortmpY(mp); ymd.ConvDateNow() == ymd.ConvDate_ymd(ymd.SortmpY(mp)) {
			key = len(mp[y][m][d]) + 1

			for _, val := range e {
				ymd.ValInMap(mp, ymd.ConvDateNow(), int(key), val)
			}
			//	fmt.Println("data!=0")
		}
	} else {

		for _, val := range e {
			ymd.ValInMap(mp, ymd.ConvDateNow(), int(key), val)
		}
		//	fmt.Println("data==0")
	}

	f, _ := os.OpenFile("mealtake.json", os.O_RDWR, 0666)
	defer f.Close()

	data, _ = json.MarshalIndent(mp, "", "	")
	f.Write(data)

}
