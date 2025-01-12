package main

import "os"

// "githubcom/DmKorshenkov/helper/bot/in"
type year map[int]int

func main() {
	os.Chdir("./DataBase")
	os.OpenFile("tmp.json", os.O_CREATE|os.O_RDWR, 0666)
}
