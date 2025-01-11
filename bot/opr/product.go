package opr

import (
	"encoding/json"
	"os"

	. "github.com/DmKorshenkov/helper/bot/ev"
)

func RemProd(prod ...Meal) {

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

func MemProd(name string) *EnergyValue {
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

func MemAllProd() []Meal {
	var tmp = make(map[string]EnergyValue)
	data, _ := os.ReadFile("product.json")
	//	fmt.Println(string(data))
	json.Unmarshal(data, &tmp)
	var slProd = make([]Meal, 0, len(tmp))

	for name, ev := range tmp {
		slProd = append(slProd, *SetO(name, ev))
	}
	return slProd
}
