package o

import (
	"fmt"
	"math"
)

type Ev struct {
	P   float64 `json:"Prot"`
	F   float64 `json:"Fat"`
	C   float64 `json:"Carb"`
	Fb  float64 `json:"Fiber,omitempty"`
	Cal float64 `json:"Cal"`
	W   Weight  `json:"W,omitempty"`
}

func NewEv() *Ev {
	return &Ev{}
}
func SetEv(p float64, f float64, c float64, fb float64) *Ev {
	return &Ev{P: p, F: f, C: c, Fb: fb, Cal: p*4 + f*9 + c*4 + fb*1.2, W: *SetW(100, "")}
}
func (Ev *Ev) SetEv(p float64, f float64, c float64, fb float64) {
	Ev.P = p
	Ev.F = f
	Ev.C = c
	Ev.Fb = fb
	Ev.Cal = p*4 + f*9 + c*4 + fb*1.2
	Ev.W.SetWeight(100)
}

func (ev *Ev) SetWeight(weight float64) *Ev {
	ev.W.Weight = weight
	return ev
}
func (Ev *Ev) SetOneGram() *Ev {
	Ev.W.Weight = 100
	Ev.P = ((Ev.P) / Ev.W.Weight)
	Ev.F = ((Ev.F) / Ev.W.Weight)
	Ev.C = ((Ev.C) / Ev.W.Weight)
	Ev.Fb = ((Ev.Fb) / Ev.W.Weight)
	Ev.Cal = ((Ev.Cal) / Ev.W.Weight)
	Ev.W.Weight = 1
	//Ev.Round()
	return Ev
}
func (Ev *Ev) SetPortion(weight float64) *Ev {

	Ev.W.Weight = weight
	Ev.P = (Ev.P) * Ev.W.Weight
	Ev.F = (Ev.F) * Ev.W.Weight
	Ev.C = (Ev.C) * Ev.W.Weight
	Ev.Fb = (Ev.Fb) * Ev.W.Weight
	Ev.Cal = (Ev.Cal) * Ev.W.Weight
	Ev.Round()
	return Ev
}
func (Ev *Ev) Round() *Ev {
	//Round округлить
	Ev.P = math.Round((Ev.P)*1000) / 1000
	Ev.F = math.Round((Ev.F)*1000) / 1000
	Ev.C = math.Round((Ev.C)*1000) / 1000
	Ev.Fb = math.Round((Ev.Fb)*1000) / 1000
	Ev.Cal = math.Round((Ev.Cal)*1000) / 1000
	return Ev
}
func (Ev *Ev) Str() string {
	return fmt.Sprintf("prot - %f\nfats - %f\ncarb - %f\nfiber - %f\ncalories - %f\n%s\n", Ev.P, Ev.F, Ev.C, Ev.Fb, Ev.Cal, Ev.W.Str())
}
func (Ev *Ev) SumEv(Ev2 Ev) Ev {
	Ev.P += Ev2.P
	Ev.F += Ev2.F
	Ev.C += Ev2.C
	Ev.Fb += Ev2.Fb
	Ev.Cal += Ev2.Cal
	Ev3 := *Ev
	return Ev3
}
func (Ev *Ev) DiffEv(Ev2 Ev) Ev {
	Ev.P -= Ev2.P
	Ev.F -= Ev2.F
	Ev.C -= Ev2.C
	Ev.Fb -= Ev2.Fb
	Ev.Cal -= Ev2.Cal
	Ev3 := *Ev
	return Ev3
}

func (Ev *Ev) DivEv() Ev {
	Ev.P /= Ev.W.Weight
	Ev.F /= Ev.W.Weight
	Ev.C /= Ev.W.Weight
	if Ev.Fb != 0 {
		Ev.Fb /= Ev.W.Weight
	}
	Ev.Cal /= Ev.W.Weight
	Ev3 := *Ev
	return Ev3
}
