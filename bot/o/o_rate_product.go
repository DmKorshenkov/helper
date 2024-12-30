package main

import (
	"encoding/json"
	"fmt"
	"os"

	. "github.com/DmKorshenkov/helper/bot/ev"
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

func RemProd(name string) EnergyValue {
	var tmp = make(map[string]EnergyValue)
	data, _ := os.ReadFile("product.json")
	json.Unmarshal(data, &tmp)
	return tmp[name]
}

func MemRate(rate Object) {
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

func RemRate() EnergyValue {
	data, _ := os.ReadFile("rate.json")
	var rate = make([]Object, 0, 12)
	json.Unmarshal(data, &rate)
	return (rate[len(rate)-1].EnergyValue)
}

func main() {
	//o := NewO()
	/*	rate := SetEv(1.2, 0.8, 3, 0).SetPortion(75)
		rate.W.SetWeight(75000)
		o := SetO("75", *rate)
		MemRate(*o)*/
	rate := RemRate()
	fmt.Println(rate)

}
