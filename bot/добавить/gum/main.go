package main

import (
	"fmt"
)

type weight struct {
	x []float64
}

type Down struct {
	count uint
	ex    []ex
}

type ex struct {
	name   string
	weight weight
	rep    []uint
}

func new_ex() *ex {
	return &ex{weight: weight{}}
}

// ///////
func (ex *ex) p() string {
	var str string
	if ex.name != "" {
		fmt.Println(ex.name)
	}
	if ex.weight.x != nil {
		for i, v := range ex.weight.x {
			if len(ex.rep) == len(ex.weight.x) {
				fmt.Printf("подход %-2v: вес - %-5v на %-3v повторений\n", i, v, ex.rep[i])
				str += fmt.Sprintf("подход %-2v: вес - %-5v на %-3v повторений\n", i, v, ex.rep[i])
			}
		}
	}
	return str
}

func (ex *ex) ex_name(name string) {
	ex.name = name
}

func (ex *ex) ex_weight(weight ...float64) {
	if (ex.weight.x) == nil {
		tmp_weight := make([]float64, 0, 5)
		tmp_weight = append(tmp_weight, weight...)
		ex.weight.x = tmp_weight
		return
	}
	ex.weight.x = append(ex.weight.x, weight...)
}

func (ex *ex) ex_rep(rep ...uint) {
	if ex.rep == nil {
		var tmp_rep = make([]uint, 0, 35)
		tmp_rep = append(tmp_rep, rep...)
		ex.rep = tmp_rep
		return
	}
	fmt.Println("!")
	ex.rep = append(ex.rep, rep...)
}

type up struct {
	ex []ex
}

func new_up() up {
	return up{}
}

func (up *up) up_ex(data_ex ...ex) {
	if up.ex == nil {
		fmt.Println("!")
		var tmp = make([]ex, 0, 9)
		up.ex = tmp
	}
	up.ex = append(up.ex, data_ex...)
}

func (up up) p() {
	for _, ex := range up.ex {
		ex.p()
	}
}

func main() {
	var ch = make(chan ex)
	var up = new_up()

	go func() {
		var ex_1 = new_ex()
		ex_1.ex_name("Жим")

		ex_1.ex_weight(100, 50, 60, 65, 67.5, 67.5)

		ex_1.ex_rep(0, 10, 8, 8, 6, 6)

		ch <- *ex_1
		close(ch)

	}()
	for {
		if val, ok := <-ch; ok {
			up.up_ex(val)
			fmt.Println(up)
			//	close(ch)
			break
		}
	}
	// up.p()
}
