package ev

import (
	"fmt"
	"math"

	. "github.com/DmKorshenkov/helper/bot/weight"
)

type EnergyValue struct {
	P   float64 `json:"Prot,omitempty"`
	F   float64 `json:"Fat,omitempty"`
	C   float64 `json:"Carb,omitempty"`
	Fb  float64 `json:"Fiber,omitempty"`
	Cal float64 `json:"Cal,omitempty"`
	W   Weight  `json:"W,omitempty"`
}

func NewEv() *EnergyValue {
	return &EnergyValue{}
}
func SetEv(p float64, f float64, c float64, fb float64) *EnergyValue {
	return &EnergyValue{P: p, F: f, C: c, Fb: fb, Cal: p*4 + f*9 + c*4 + fb*1.2, W: *SetW(100, "")}
}
func (ev *EnergyValue) SetEv(p float64, f float64, c float64, fb float64) {
	ev.P = p
	ev.F = f
	ev.C = c
	ev.Fb = fb
	ev.Cal = p*4 + f*9 + c*4 + fb*1.2
	ev.W.SetWeight(100)
}
func (ev *EnergyValue) SetOneGram() *EnergyValue {
	ev.W.Weight = 100
	ev.P = ((ev.P) / ev.W.Weight)
	ev.F = ((ev.F) / ev.W.Weight)
	ev.C = ((ev.C) / ev.W.Weight)
	ev.Fb = ((ev.Fb) / ev.W.Weight)
	ev.Cal = ((ev.Cal) / ev.W.Weight)
	ev.W.Weight = 1
	//ev.Round()
	return ev
}
func (ev *EnergyValue) SetPortion(weight float64) *EnergyValue {

	ev.W.Weight = weight
	ev.P = (ev.P) * ev.W.Weight
	ev.F = (ev.F) * ev.W.Weight
	ev.C = (ev.C) * ev.W.Weight
	ev.Fb = (ev.Fb) * ev.W.Weight
	ev.Cal = (ev.Cal) * ev.W.Weight
	ev.Round()
	return ev
}
func (ev *EnergyValue) Round() *EnergyValue {
	//Round округлить
	ev.P = math.Round((ev.P)*1000) / 1000
	ev.F = math.Round((ev.F)*1000) / 1000
	ev.C = math.Round((ev.C)*1000) / 1000
	ev.Fb = math.Round((ev.Fb)*1000) / 1000
	ev.Cal = math.Round((ev.Cal)*1000) / 1000
	return ev
}
func (ev *EnergyValue) Str() string {
	return fmt.Sprintf("prot - %f\nfats - %f\ncarb - %f\nfiber - %f\ncalories - %f\n%s\n", ev.P, ev.F, ev.C, ev.Fb, ev.Cal, ev.W.Str())
}
func (ev EnergyValue) SumEv(ev2 EnergyValue) EnergyValue {
	ev.P += ev2.P
	ev.F += ev2.F
	ev.C += ev2.C
	ev.Fb += ev2.Fb
	ev.Cal += ev2.Cal
	return ev
}
func (ev *EnergyValue) DiffEv(ev2 EnergyValue) EnergyValue {
	ev.P -= ev2.P
	ev.F -= ev2.F
	ev.C -= ev2.C
	ev.Fb -= ev2.Fb
	ev.Cal -= ev2.Cal
	return *ev
}
