package ev

import (
	"fmt"
	"math"
)

type EnergyValue struct {
	Weight float64 `json:"weight,omitempty"`
	P      float64 `json:"proteins,omitempty"`
	F      float64 `json:"fats,omitempty"`
	Crb    float64 `json:"carbohydrate,omitempty"`
	Fiber  float64 `json:"fiber,omitempty"`
	Cal    float64 `json:"calories"`
}

func NewEv(get ...float64) *EnergyValue {
	var ev = new(EnergyValue)
	ev.P = get[0]
	ev.F = get[1]
	ev.Crb = get[2]
	if len(get) == 3 || len(get) == 4 {
		ev.Weight = 100
	}
	if len(get) == 4 {
		ev.Fiber = get[3]
	}
	if len(get) == 5 {
		ev.Weight = get[4]
	}
	ev.SetCal()
	//	ev.SetWeight(100)
	//hello
	return ev
}

func SumEv(ev1 EnergyValue, ev2 EnergyValue) EnergyValue {
	//var sum =
	return EnergyValue{
		P:     ev1.P + ev2.P,
		F:     ev1.F + ev2.F,
		Crb:   ev1.Crb + ev2.Crb,
		Fiber: ev1.Fiber + ev2.Fiber,
		Cal:   ev1.Cal + ev2.Cal,
	}
}

func DiffEv(ev1 EnergyValue, ev2 EnergyValue) EnergyValue {
	//var sum =
	return EnergyValue{
		P:     ev1.P - ev2.P,
		F:     ev1.F - ev2.F,
		Crb:   ev1.Crb - ev2.Crb,
		Fiber: ev1.Fiber - ev2.Fiber,
		Cal:   ev1.Cal - ev2.Cal,
	}
}

func (ev *EnergyValue) Str() string {
	str := fmt.Sprintf("Вес - %v\nКоличество белков - %.3f\nКоличество жиров - %.3f\nКоличество углеводов - %.3f\nКоличество клетчатки - %v\nКалорийность - %v\n", ev.Weight, ev.P, ev.F, ev.Crb, ev.Fiber, ev.Cal)
	return str
}

func (ev *EnergyValue) SetWeight(n float64) {
	ev.Weight = n
}

func (ev *EnergyValue) SetOneGram() *EnergyValue {
	ev.P = ev.P / ev.Weight
	ev.F = ev.F / ev.Weight
	ev.Crb = ev.Crb / ev.Weight
	ev.Fiber = ev.Fiber / ev.Weight
	ev.Weight = 1
	ev.SetCal()
	return ev
}

func (ev *EnergyValue) SetPortion(weight float64) *EnergyValue {
	//ev.SetPFC()
	ev.Weight = weight
	ev.P = math.Round(((ev.P)*ev.Weight)*100) / 100
	ev.F = math.Round(((ev.F)*ev.Weight)*100) / 100
	ev.Crb = math.Round(((ev.Crb)*ev.Weight)*100) / 100
	ev.Fiber = math.Round(((ev.Fiber)*ev.Weight)*100) / 100
	ev.SetCal()
	return ev
}
func (ev *EnergyValue) SetCal() {
	ev.Cal = math.Round((ev.P*4+ev.F*9+ev.Crb*4)*100) / 100
}
