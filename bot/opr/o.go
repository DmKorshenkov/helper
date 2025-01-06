package opr

import (
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

func (o *Object) SetO(name string, ev EnergyValue) {
	o.Name = name
	o.EnergyValue = ev
}

func (o *Object) SetName(name string) {
	o.Name = name
}

func (o *Object) SetEv(ev EnergyValue) {
	o.EnergyValue = ev
}

func (o *Object) A_weight(weight float64) *Object {
	//fmt.Println(o.EnergyValue.W.Weight, "- before")
	o.EnergyValue.W.Weight = weight
	//fmt.Println(o.EnergyValue.W.Weight, "- after")
	return o
}
