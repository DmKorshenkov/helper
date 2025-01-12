package o

import (
	"encoding/json"
	"fmt"
	"os"

	. "github.com/DmKorshenkov/helper/bot/ev"
	. "github.com/DmKorshenkov/helper/bot/ymd"
)

type Object struct {
	Name        string      `json:"Name"`
	EnergyValue EnergyValue `json:"Energy Value"`
}

func NewO() *Object {
	return &Object{}
}

func SetO(name string, ev EnergyValue) *Object {
	return &Object{Name: name, EnergyValue: ev}
}

func (o *Object) SetO(name string, ev EnergyValue) {
	o.Name = name
	o.EnergyValue = ev
}

func MemProd(prod ...Object) {

	f, _ := os.OpenFile("product.json", os.O_CREATE|os.O_RDWR, 0666)
	data, _ := os.ReadFile(f.Name())
	var tmp = make(map[string]EnergyValue)
	if len(data) != 0 {
		json.Unmarshal(data, &tmp)
	}

	for _, element := range prod {
		tmp[element.Name] = element.EnergyValue
	}
	data, _ = json.MarshalIndent(tmp, "", "	")
	f.Write(data)
	f.Close()
}

func RemProd(name string) *EnergyValue {
	var tmp = make(map[string]EnergyValue)
	data, _ := os.ReadFile("product.json")
	//	fmt.Println(string(data))
	json.Unmarshal(data, &tmp)
	//	fmt.Println(tmp[name])
	if val, ok := tmp[name]; ok {
		return &val
	} else {
		return nil
	}
}

func SetRate(rate Object) {
	rate2 := rate.EnergyValue
	defer MemRate(rate2)
	f, _ := os.OpenFile("rate.json", os.O_CREATE|os.O_RDWR, 0666)
	data, _ := os.ReadFile(f.Name())
	tmp := make([]Object, 0, 12)
	if len(data) != 0 {
		json.Unmarshal(data, &tmp)
	}
	tmp = append(tmp, rate)
	data, _ = json.MarshalIndent(tmp, "", "	")
	f.Write(data)
	f.Close()
}

func MemRate(rate EnergyValue) {
	f, _ := os.OpenFile("ratetmp.json", os.O_CREATE|os.O_RDWR, 0666)
	data, _ := os.ReadFile(f.Name())
	if len(data) != 0 {
		json.Unmarshal(data, &EnergyValue{})
	}
	tmp := rate.Round()
	data, _ = json.MarshalIndent(tmp, "", "		")
	f.Write(data)
	f.Close()
}

func RemRate() *EnergyValue {
	data, _ := os.ReadFile("ratetmp.json")
	rate := EnergyValue{}
	json.Unmarshal(data, &rate)
	return &rate
}

type req struct {
	name   string
	weight float64
}

func set(n string, w float64) req {
	return req{name: n, weight: w}
}

/*func main() {

/*rate := SetEv(1.2, 0.8, 3, 0).SetPortion(75)
o := SetO("75 kg", *rate)
SetRate(*o)*/
//	o := SetO("курица", EnergyValue{W: weight.Weight{Weight: 50}})
//	MemRate(RemRate().DiffEv(EnergyValue{P: 1, F: 2, C: 3, Fb: 0}))

//
// ymd := sl.CmapYear[int, int](sl.ConvDataOne() + 1)
// fmt.Println(ymd)
//
/*
	var sl2 = make([]Object, 0, 5)
	sl2 = append(sl2, *SetO("геркулес", EnergyValue{W: weight.Weight{Weight: 40}}))
	sl2 = append(sl2, *SetO("яйца", EnergyValue{W: weight.Weight{Weight: 120}}))
	mealTake(sl2)
	sl2 = nil
	sl2 = append(sl2, *SetO("курица", EnergyValue{W: weight.Weight{Weight: 50}}))
	sl2 = append(sl2, *SetO("гречка", EnergyValue{W: weight.Weight{Weight: 60}}))
	mealTake(sl2)
	sl2 = nil
	sl2 = append(sl2, *SetO("курица", EnergyValue{W: weight.Weight{Weight: 100}}))
	sl2 = append(sl2, *SetO("гречка", EnergyValue{W: weight.Weight{Weight: 80}}))
	mealTake(sl2)
	sl2 = nil
	sl2 = append(sl2, *SetO("курица", EnergyValue{W: weight.Weight{Weight: 100}}))
	sl2 = append(sl2, *SetO("гречка", EnergyValue{W: weight.Weight{Weight: 80}}))
	mealTake(sl2)
	_ = sl2*/ /*
	ev := RemRate()
	ev.DiffEv(*ev)
	MemRate(*ev)

}*/

func mealTake(o2 []Object) {
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
	sumev := EnergyValue{}
	for _, food := range o {
		sumev.SumEv(food.EnergyValue)
	}
	o = append(o, *SetO("БЖУ", sumev))
	MemMeal(o)

	MemRate(RemRate().DiffEv(o[len(o)-1].EnergyValue))
}

func meal(food Object) *Object {
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

		if y, m, d := SortmpY(mp); ConvDateNow() == ConvDate_ymd(SortmpY(mp)) {
			key = len(mp[y][m][d]) + 1

			for _, val := range e {
				ValInMap(mp, ConvDateNow(), int(key), val)
			}
			//	fmt.Println("data!=0")
		}
	} else {

		for _, val := range e {
			ValInMap(mp, ConvDateNow(), int(key), val)
		}
		//	fmt.Println("data==0")
	}

	f, _ := os.OpenFile("mealtake.json", os.O_RDWR, 0666)
	defer f.Close()

	data, _ = json.MarshalIndent(mp, "", "	")
	f.Write(data)

}
