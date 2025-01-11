package opr

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/DmKorshenkov/helper/bot/ev"
	"github.com/DmKorshenkov/helper/bot/ymd"
)

type Prod struct {
	Name   string
	Weight float64
}

func MemMealTake(o []Prod) {
	var meal_take = make([]Meal, 0, 10)
	for in, food := range o {
		//	fmt.Println(food)
		if check := meal(food); check != nil {
			meal_take[in] = *check
		} else {
			fmt.Println("check==nil")
		}

	}
	//все продукты в o
	//fmt.Println(o)
	sumev := ev.EnergyValue{}
	for _, food := range meal_take {
		sumev.SumEv(food.EnergyValue)
	}
	meal_take = append(meal_take, *SetO("сьел", sumev))

	rate := func() ev.EnergyValue {
		rate := RemRate()
		rate.DiffEv(sumev)
		MemRate(*rate)
		return *rate
	}()
	meal_take = append(meal_take, *SetO("осталось", rate))
	memMeal(meal_take)

	//MemRate(RemRate().DiffEv(o[len(o)-1].EnergyValue))
}

func meal(food Prod) *Meal {
	//ищет продукт+возвращает его порцию
	if f := MemProd(food.Name); f != nil {
		f.SetOneGram().SetPortion(food.Weight)
		return (SetO(food.Name, *f))
	} else {
		fmt.Println("f==nil")
		return nil
	}

}

func memMeal(e []Meal) {

	var mp = make(map[int]map[int]map[int]map[int][]Meal)
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

	f, _ := os.OpenFile("mealtake.json", os.O_CREATE|os.O_RDWR, 0666)

	data, _ = json.MarshalIndent(mp, "", "	")
	_, _ = f.Write(data)
	_ = f.Close()

}
