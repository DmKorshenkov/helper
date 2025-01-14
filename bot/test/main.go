package main

import (
	"os"

	"github.com/DmKorshenkov/helper/bot/check"
	"github.com/DmKorshenkov/helper/bot/o"
)

// "githubcom/DmKorshenkov/helper/bot/in"

type year map[int]int

func main() {
	//	os.OpenFile("tmp.json", os.O_CREATE|os.O_RDWR, 0666)
	os.Chdir(".././DataBase")
	p := check.RemFood("рыба 100 100 100")
	if len(p) != 0 {
		o.RemFood(p...)
	} else {
	}
}
