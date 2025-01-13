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
	var weight = check.RemWeight("77.1 morning")
	if weight != nil {
		fmt.Println(weight)
	} else {
		fmt.Println("weight==false")
	}
	//check Rate
	var rate = check.Rate("1.2 0.8 3 75")
	if rate != nil {
		fmt.Println(rate)
		//o.RemRate(*rate)
	} else {
		fmt.Println("rate==false")
	}

	//Check Food Ev(product)
	var food = check.RemFood("курица 20 4 0\nгречка 0 1")
	if len(food) != 0 {
		fmt.Println(food)
	} else {
		fmt.Println("food==false")
	}

	//check Prod Meal Take
	var prod = check.Prod("курица 100\nxyi 100")
	if len(prod) != 0 {
		fmt.Println(prod)
	} else {
		fmt.Println("prod==false")
	}

}
