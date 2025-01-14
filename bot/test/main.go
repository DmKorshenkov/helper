package main

import (
	"fmt"
	"os"

	"github.com/DmKorshenkov/helper/bot/check"
)

// "githubcom/DmKorshenkov/helper/bot/in"

type year map[int]int

func main() {
	//	os.OpenFile("tmp.json", os.O_CREATE|os.O_RDWR, 0666)
	os.Chdir(".././DataBase")
	//check Weight
	var weight = check.RemWeight("77.1m morning")
	fmt.Println(weight)

	var food = check.RemFood("курица 21 3 0")
	fmt.Println(food)

}
