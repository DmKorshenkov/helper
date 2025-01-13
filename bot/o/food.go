package o

import (
	"encoding/json"
	"os"
)

type Food struct {
	Name        string `json:"Name"`
	EnergyValue Ev     `json:"Energy Value"`
}

func NewFood() *Food {
	return &Food{}
}

func SetFood(name string, EnergyValue Ev) *Food {
	return &Food{Name: name, EnergyValue: EnergyValue}
}

func (o *Food) SetFood(name string, EnergyValue Ev) {
	o.Name = name
	o.EnergyValue = EnergyValue
}

func (o *Food) SetName(name string) {
	o.Name = name
}

func (o *Food) SetEnergyValue(EnergyValue Ev) {
	o.EnergyValue = EnergyValue
}

func (o *Food) Food_weight(weight float64) *Food {
	//fmt.Println(o.Ev.W.Weight, "- before")
	o.EnergyValue.W.Weight = weight
	//fmt.Println(o.Ev.W.Weight, "- after")
	return o
}

func RemFood(prod ...Food) {

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

func MemFood(name string) *Ev {
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

func MemAllFood() []Food {
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
