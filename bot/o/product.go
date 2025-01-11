package o

import (
	"encoding/json"
	"os"
)

func RemProd(prod ...Food) {

	f, _ := os.OpenFile("product.json", os.O_CREATE|os.O_RDWR, 0666)
	data, _ := os.ReadFile(f.Name())
	var tmp = make(map[string]Ev)
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

func MemProd(name string) *Ev {
	var tmp = make(map[string]Ev)
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

func MemAllProd() []Food {
	var tmp = make(map[string]Ev)
	data, _ := os.ReadFile("product.json")
	//	fmt.Println(string(data))
	json.Unmarshal(data, &tmp)
	var slProd = make([]Food, 0, len(tmp))

	for name, EnergyValue := range tmp {
		slProd = append(slProd, *SetFood(name, EnergyValue))
	}
	return slProd
}
