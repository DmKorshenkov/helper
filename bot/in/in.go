package in

import (
	"fmt"
	"strings"

	"github.com/DmKorshenkov/helper/bot/check"
	"github.com/DmKorshenkov/helper/bot/mr"
)

type I struct {
	cmd  string
	key  string
	req  uint8
	data string
}

func NewI() *I {
	return &I{}
}

func (i *I) NewI(cmd string, key string) *I {
	//i.cmd = strings.TrimSpace(cmd)
	i.cmd = strings.ToLower(strings.TrimSpace(cmd))
	i.key = helpTrim(key)
	//i.data = data
	//i.key = strings.TrimSpace(key)
	return i
}

func (i *I) CheckCmd() {
	switch i.cmd {
	case "rem":
		i.req = 10
		return
	case "mem":
		i.req = 20
		return
	case "cal":
		i.req = 30
		return
	case "запомни":
		i.req = 10
		return
	case "вспомни":
		i.req = 20
		return
	case "калькулятор":
		i.req = 30
		return
	default:
		return
	}
}

func (i *I) CheckKey() {
	switch i.key {
	case "weight":
		i.req += 1
		return
	case "prod":
		i.req += 2
		return
	case "meal take":
		i.req += 3
		return
	case "вес":
		i.req += 1
		return
	case "продукт":
		i.req += 2
		return
	case "прием пищи":
		i.req += 3
		return
	default:
		i.req = 0
		return
	}
}

func (i *I) Check() {
	i.CheckCmd()
	i.CheckKey()
}

func (i *I) PI() {
	fmt.Println("input:")
	fmt.Println("i.cmd = ", i.cmd)
	fmt.Println("i.key = ", i.key)
}

func In(msg string) {
	var i = NewI()
	var mr mr.RemMem

	message := strings.SplitN(msg, "\n", 2)
	cmdkey := strings.SplitN(message[0], " ", 2)
	i.NewI(cmdkey[0], cmdkey[1])
	i.PI()
	i.Check()
	if i.req == 0 {
		fmt.Println("cmd/key!=true")
		return
	} else {
		fmt.Println("cmd/key==true")
	}

	if len(message) > 1 {
		i.data = message[1]
	}
	func(i I) {
		switch i.req {
		case 11:
			//RemWeight()
			fmt.Println(i.req)
			fmt.Println()
			//проверит data и создаст weight
			w := check.CheckRemWeight(i.data)
			if w == nil {
				fmt.Println("CheckRemWeight == nil")
				//check == false
				//return
			} else {
				mr.Rem(w)
				fmt.Println(w)
				//Rem()
			}
		case 21:
			//MemWeight()
			fmt.Println(i.req)
			fmt.Println()
			//Check
			//Mem()
		case 12:
			//RemProd()
			fmt.Println(i.req)
			fmt.Println()
			meal := check.CheckRemProd(i.data)
			if meal == nil {
				fmt.Println("CheckRemProd == nil")
				//return
			} else {
				fmt.Println(meal)
				//Rem()
			}
			//Check
		case 22:
			//MemMealTake
			fmt.Println(i.req)
			//Check
			//Mem()
		case 13:
			fmt.Println(i.req)
			fmt.Println()
			meal_take := check.CheckRemProd(i.data)
			if meal_take == nil {
				fmt.Println("CheckRemMl == nil")
				//return nil
			} else {
				fmt.Println(meal_take)
				//Rem()
			}
			//Check
		case 23:
			fmt.Println(i.req)
			//Check
			//Mem()
		}
	}(*i)
}

func helpTrim(str string) string {
	return strings.ToLower(strings.TrimSuffix(strings.TrimSuffix(strings.TrimPrefix(str, " "), "\n"), " "))
}
