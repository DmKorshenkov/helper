package fnc

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/DmKorshenkov/helper/bot/o"
	"github.com/DmKorshenkov/helper/bot/ymd"
)

func MealTake(p ...o.Prod) []o.Food {

	var mealTake = make([]o.Food, 0, len(p))
	var sumev = o.Ev{}
	//
	for _, food := range p {
		//
		mealTake = append(mealTake, getFood(food))
	}
	//
	//
	for _, food := range mealTake {
		sumev.SumEv(food.EnergyValue)
	}
	mealTake = append(mealTake, *o.SetFood("БЖУ", sumev))
	//RemRateDay
	func() {
		rate := *o.MemRate()
		rate.DiffEv(sumev)
		rate.Round()
		fmt.Println(rate)
		//o.RemRateDay(rate)
	}()
	return mealTake
}

func getFood(food o.Prod) o.Food {
	// get food
	f := o.MemFood(food.Name)
	f.SetOneGram().SetPortion(food.Weight)
	return *(o.SetFood(food.Name, *f))

}

func RemMeal(e []o.Food) {
	f, _ := os.OpenFile("mealTake.json", os.O_CREATE|os.O_RDWR, 0666)
	var mp = make(map[int]map[int]map[int]map[int][]o.Food)
	key := 1
	data, _ := os.ReadFile("mealTake.json")
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

	data, _ = json.MarshalIndent(mp, "", "	")
	_, _ = f.Write(data)
	_ = f.Close()
}
