package food

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

type EnergyValue struct {
	Weight float64 `json:"weight"`
	P      float64 `json:"proteins,omitempty"`
	F      float64 `json:"fats,omitempty"`
	Crb    float64 `json:"carbohydrate,omitempty"`
	Fiber  float64 `json:"fiber,omitempty"`
	Cal    float64 `json:"calories"`
}

func Energy_Value(get ...string) EnergyValue {
	var food = new(EnergyValue)
	if len(get) == 3 {
		food.P = help(get[0])
		food.F = help(get[1])
		food.Crb = help(get[2])
		food.Fiber = help("0")
	} else if len(get) == 4 {
		food.P = help(get[0])
		food.F = help(get[1])
		food.Crb = help(get[2])
		food.Fiber = help(get[3])
	}
	food.SetCal()
	food.SetWeight(100)
	//hello
	return (*food)
}

func (ev *EnergyValue) PrintEv() string {
	str := fmt.Sprintf("Вес - %v\nКоличество белков - %.1f\nКоличество жиров - %.1f\nКоличество углеводов - %.1f\nКоличество клетчатки - %v\nКалорийность - %v\n", ev.Weight, ev.P, ev.F, ev.Crb, ev.Fiber, ev.Cal)
	return str
}

func (ev *EnergyValue) SetWeight(n float64) {
	ev.Weight = n
}

func (ev *EnergyValue) SetOneGram() {
	ev.P = ev.P / ev.Weight
	ev.F = ev.F / ev.Weight
	ev.Crb = ev.Crb / ev.Weight
	ev.Fiber = ev.Fiber / ev.Weight
}

func (ev *EnergyValue) SetCal() {
	ev.Cal = math.Round((ev.P*4+ev.F*9+ev.Crb*4)*100) / 100
}
func (ev *EnergyValue) SetPortion(weight float64) {
	ev.SetWeight(weight)
	//ev.SetPFC()
	ev.P *= ev.Weight / 100
	ev.F *= ev.Weight / 100
	ev.Crb *= ev.Weight / 100
	ev.Fiber *= ev.Weight / 100
	ev.SetCal()
}
func help(get string) (f64 float64) {
	f64, err := strconv.ParseFloat(strings.TrimSpace(get), 64)
	if err != nil {
		log.Println(err.Error())
	}
	//	f64 = f64 * 100 / 100
	f64 = math.Round(f64*100) / 100
	return f64
}
