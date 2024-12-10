package food

import (
	"fmt"
	"log"
	"math"
)

type EnergyValue struct {
	W  float64 `json:"weight"`
	P  float64 `json:"protein"`
	F  float64 `json:"fat"`
	C  float64 `json:"Carbohydrate"`
	Cl float64 `json:"Calories"`
}

func SetCal(p, f, c float64) float64 {
	x := ((p * 4) + (f * 9) + (c * 4))
	return math.Round(x*100) / 100
}

// исправить функцию SetCal
func (food *EnergyValue) SetCal2() {
	food.Cl = math.Round(((food.P*4)+(food.F*9)+(food.C*4))*100) / 100
}

func (ev *EnergyValue) PFC100(p ...float64) {
	if len(p) != 3 {
		log.Println("len(PFC100) != 3 - error")
	}
	ev.W = 100
	ev.P = p[0] / ev.W
	ev.F = p[1] / ev.W
	ev.C = p[2] / ev.W
	ev.Cl = math.Round(SetCal(ev.P, ev.F, ev.C)*100) / 100

}

func (ev *EnergyValue) PortionEatUncooked(weight float64) {
	ev.W = weight
	ev.P = math.Round((ev.P * ev.W))
	ev.F = math.Round((ev.F * ev.W))
	ev.C = math.Round((ev.C * ev.W))
	ev.Cl = SetCal(ev.P, ev.F, ev.C)
}

func (ev *EnergyValue) PortionCooked(weight float64) {
	ev.W = weight
	ev.P = math.Round((ev.P / ev.W * 1000)) / 1000
	ev.F = math.Round((ev.F / ev.W * 1000)) / 1000
	ev.C = math.Round((ev.C / ev.W * 1000)) / 1000
	ev.Cl = SetCal(ev.P, ev.F, ev.C)
}

func (ev *EnergyValue) PortionEatCooked(weight float64) {
	ev.W = weight
	multiplier := func(x, y float64) float64 { return math.Round(x*y*100) / 100 }
	ev.P = multiplier(ev.P, ev.W)
	ev.F = multiplier(ev.F, ev.W)
	ev.C = multiplier(ev.C, ev.W)
	ev.Cl = SetCal(ev.P, ev.F, ev.C)

}

func main() {
	chicken := new(EnergyValue)
	//	Type(chicken)
	chicken.PFC100(7, 1, 79)
	fmt.Printf("БЖУ продукта на 1 грамм - %v\n\n", chicken)
	//Type(chicken)
	chicken.PortionEatUncooked(50)
	fmt.Printf("БЖУ на порцию и вес порции  - %v\n\n", chicken)
	//Type(chicken)
	chicken.PortionCooked(300)
	fmt.Printf("В приготовленной порции продукта БЖУ и вес - %v\n\n", chicken)
	chicken.PortionEatCooked(300)
	fmt.Printf("В порции сьедено - %v", chicken)
	//	Type(chicken)

}

func Type(t any) {
	fmt.Printf("%T\n%v\n", t, t)
}

func check(err error) {
	if err != nil {
		log.Println(err)
	}
}
