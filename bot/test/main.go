package main

import (
	"bufio"
	"os"

	"github.com/DmKorshenkov/helper/bot/in"
)

func main() {
	os.Chdir(".././opr")
	//	os.Chdir(".././DataBase")
	/*msg := os.Args
	if len(msg) > 1 {
		In(msg[1:]...)
	}*/
	r := bufio.NewReader(os.Stdin)
	i, _ := r.ReadString('\n')
	in.In(i)
	// In(msg[1:])
}
