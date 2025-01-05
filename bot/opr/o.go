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
