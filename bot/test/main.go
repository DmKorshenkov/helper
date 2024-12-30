package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"time"
)

type Weight struct {
	Info   string  `json:"info,omitempty"`
	Weight float64 `json:",omitempty"`
}

type EnergyValue struct {
	P   float64 `json:"Prot,omitempty"`
	F   float64 `json:"Fat,omitempty"`
	C   float64 `json:"Carb,omitempty"`
	Fb  float64 `json:"Fiber,omitempty"`
	Cal float64 `json:"Cal,omitempty"`
	W   Weight  `json:"W,omitempty"`
}

type Object struct {
	Name        string      `json:"Name"`
	EnergyValue EnergyValue `json:"Energy Value"`
}

// Object
func NewO() *Object {
	return &Object{}
}
func NewOSetName(name string) *Object {
	return &Object{Name: name}
}
func (o *Object) SetName(name string) {
	o.Name = name
}

// Weight
func NewW(w float64) *Weight {
	return &Weight{Weight: w}
}
func (w *Weight) SetWeight(weight float64) {
	w.Weight = weight
}
func (w *Weight) SetInfo(info string) {
	w.Info = info
}

// Energy Value
func NewEv(p float64, f float64, c float64, fb float64) *EnergyValue {
	return &EnergyValue{}
}
func (ev *EnergyValue) SetEv(p float64, f float64, c float64, fb float64) {
	ev.P = p
	ev.F = f
	ev.C = c
	ev.Fb = fb
	ev.Cal = p*4 + f*9 + c*4 + fb*1.2
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

// Round - округлить
func (ev *EnergyValue) Round() *EnergyValue {
	//Round округлить
	ev.P = math.Round((ev.P)*1000) / 1000
	ev.F = math.Round((ev.F)*1000) / 1000
	ev.C = math.Round((ev.C)*1000) / 1000
	ev.Fb = math.Round((ev.Fb)*1000) / 1000
	ev.Cal = math.Round((ev.Cal)*1000) / 1000
	return ev
}

// Sum
func (ev EnergyValue) SumEv(ev2 EnergyValue) EnergyValue {
	ev.P += ev2.P
	ev.F += ev2.F
	ev.C += ev2.C
	ev.Fb += ev2.Fb
	ev.Cal += ev2.Cal
	return ev
}

// Difference
func (ev *EnergyValue) DiffEv(ev2 EnergyValue) EnergyValue {
	ev.P -= ev2.P
	ev.F -= ev2.F
	ev.C -= ev2.C
	ev.Fb -= ev2.Fb
	ev.Cal -= ev2.Cal
	return *ev
}

func main() {
	var ch = make(chan int)
	for {
		print("\nbefore\n")

		go help(ch)
		print("\nafter\n")

		for val := range ch {
			fmt.Println(val)
			if val == 60 {
				fmt.Println("close")
				close(ch)
			}
		}

		for i := 100; i < 160; i++ {
			fmt.Println(i)
		}
	}

}

func help(ch chan int) {
	var i int
	for {
		time.Sleep(time.Millisecond * 300)
		i++
		ch <- i

		//time.Sleep(time.Second * 2)
	}
}

func Eat(o *EnergyValue, eat EnergyValue) {
	//	fmt.Println(eat)
	o.DiffEv(eat)
}

func RemRate() EnergyValue {
	data, err := os.ReadFile("rate.json")
	if err != nil {
		log.Println(err.Error())
	}

	rate := NewO()
	json.Unmarshal(data, rate)
	return rate.EnergyValue
}
