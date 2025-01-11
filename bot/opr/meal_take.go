package opr

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/DmKorshenkov/helper/bot/o"
	"github.com/DmKorshenkov/helper/bot/ymd"
)

type Prod struct {
	Name   string
	Weight float64
}

func RemMealTake(p []Prod) {
	var meal_take = make([]o.Food, 0, 10)
	for in, food := range p {
		//	fmt.Println(food)
		if check := meal(food); check != nil {
			meal_take[in] = *check
		} else {
			fmt.Println("check==nil")
		}

	}
	//все продукты в o
	//fmt.Println(o)
	sumev := o.Ev{}
	for _, food := range meal_take {
		sumev.SumEv(food.EnergyValue)
	}
	meal_take = append(meal_take, *o.SetFood("сьел", sumev))

	rate := func() o.Ev {
		rate := o.MemRate()
		rate.DiffEv(sumev)
		o.RemRate(*rate)
		return *rate
	}()
	meal_take = append(meal_take, *o.SetFood("осталось", rate))
	memMeal(meal_take)

	//MemRate(RemRate().DiffEv(o[len(o)-1].EnergyValue))
}

func meal(food Prod) *o.Food {
	//ищет продукт+возвращает его порцию
	if f := o.MemProd(food.Name); f != nil {
		f.SetOneGram().SetPortion(food.Weight)
		return (o.SetFood(food.Name, *f))
	} else {
		fmt.Println("f==nil")
		return nil
	}

}

func memMeal(e []o.Food) {

	var mp = make(map[int]map[int]map[int]map[int][]o.Food)
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
