package main

import (
	"bufio"
	"os"

	"github.com/DmKorshenkov/helper/bot/in"
)

// "githubcom/DmKorshenkov/helper/bot/in"
type year map[int]int

func main() {
	os.Chdir("//opr")
	r := bufio.NewReader(os.Stdin)
	i, _ := r.ReadString('\n')
	in.In(i)

}
