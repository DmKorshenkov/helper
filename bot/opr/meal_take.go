package opr

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/DmKorshenkov/helper/bot/ev"
	"github.com/DmKorshenkov/helper/bot/ymd"
)

func MealTake(o2 []Object) {
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
	o = append(o, *SetO("сьел", sumev))

	rate := func() ev.EnergyValue {
		rate := RemRate()
		rate.DiffEv(sumev)
		MemRate(*rate)
		return *rate
	}()
	o = append(o, *SetO("осталось", rate))
	MemMeal(o)

	//MemRate(RemRate().DiffEv(o[len(o)-1].EnergyValue))
}

func meal(food Object) *Object {
	//ищет продукт+возвращает его порцию
	if f := RemProd(food.Name); f != nil {
		f.SetOneGram().SetPortion(food.EnergyValue.W.Weight)
		return (SetO(food.Name, *f))
	} else {
		fmt.Println("f==nil")
		return nil
	}

}

func MemMeal(e []Object) {

	var mp = make(map[int]map[int]map[int]map[int][]Object)
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
