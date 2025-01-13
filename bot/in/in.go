package in

import (
	"fmt"
	"strings"

	"github.com/DmKorshenkov/helper/bot/check"
)

type I struct {
	cmd  string
	key  string
	req  int32
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
		i.req = 'r'
		return
	case "mem":
		i.req = 'm'
		return
	case "cal":
		i.req = 'c'
		return
	case "запомни":
		i.req = 'r'
		return
	case "вспомни":
		i.req = 'm'
		return
	case "калькулятор":
		i.req = 'c'
		return
	default:
		return
	}
}

func (i *I) CheckKey() {
	switch i.key {
	case "weight":
		i.req += 'w'
		return
	case "prod":
		i.req += 'f'
		return
	case "rate":
		i.req += 'R'
		return
	case "meal take":
		i.req += 'm'
		return
	case "вес":
		i.req += 'w'
		return
	case "продукт":
		i.req += 'f'
		return
	case "норму бжу":
		i.req += 'R'
		return
	case "прием пищи":
		i.req += 'm'
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

func In(msg string) string {
	var i = NewI()

	checK, got := func() ([]string, []string) {
		get := strings.SplitN(msg, "\n", 2)
		cmdkey := strings.SplitN(get[0], " ", 2)
		return cmdkey, get
	}()
	if len(checK) != 2 {
		return "Упс)\nНеверное количество ключевых слов в команде.\n"
	}
	i.NewI(checK[0], checK[1])
	i.PI()
	i.Check()
	if i.req == 0 {
		return "cmd/key!=true\n"
	} else {
		fmt.Println("cmd/key==true")
	}

	if len(got) > 1 {
		i.data = got[1]
	}
	func(i I) {
		switch i.req {
		case 'r' + 'w':
			//rem weight
			fmt.Println("!")
			weight := check.RemWeight(i.data)
			_ = weight
		case 'm' + 'w':
			// mem weight

		case 'r' + 'f':
			//rem prod/food

		case 'r' + 'R':
			//rem rate

		case 'm' + 'R':
			//mem rate

		case 'r' + 'm':
			//rem meal

		case 'm' + 'm':
			//mem meal

		}
	}(*i)
	return "in.In(msg) - mission complete, bro)"
}

func helpTrim(str string) string {
	return strings.ToLower(strings.TrimSuffix(strings.TrimSuffix(strings.TrimPrefix(str, " "), "\n"), " "))
}
