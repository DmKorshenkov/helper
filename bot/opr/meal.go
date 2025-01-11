package opr

import (
	. "github.com/DmKorshenkov/helper/bot/ev"
)

type Meal struct {
	Name        string      `json:"Name"`
	EnergyValue EnergyValue `json:"Energy Value"`
}

func NewO() *Meal {
	return &Meal{}
}

func SetO(name string, ev EnergyValue) *Meal {
	return &Meal{Name: name, EnergyValue: ev}
}

func (o *Meal) SetO(name string, ev EnergyValue) {
	o.Name = name
	o.EnergyValue = ev
}

func (o *Meal) SetName(name string) {
	o.Name = name
}

func (o *Meal) SetEv(ev EnergyValue) {
	o.EnergyValue = ev
}

func (o *Meal) A_weight(weight float64) *Meal {
	//fmt.Println(o.EnergyValue.W.Weight, "- before")
	o.EnergyValue.W.Weight = weight
	//fmt.Println(o.EnergyValue.W.Weight, "- after")
	return o
}
