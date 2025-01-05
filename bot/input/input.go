package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/DmKorshenkov/helper/bot/sl"
)

type input struct {
	messageIn  []string
	lenIn      int
	MessageOut string
	object     *object
}

type object struct {
	command string
	key     string
	data    string
	status  bool
}

func (in *input) CreateRequest(message string) {
	var tmp = make([]string, 0, 2)
	tmp = strings.SplitAfterN(in.messageIn[0], " ", 2)
	if len(tmp) != 2 {
		in.messageIn = nil
		in.lenIn = 0
		in.MessageOut = "неверное количество слов"
		return
	}

	in.object.command = strings.ToLower(strings.TrimSpace(tmp[0]))
	in.object.key = strings.ToLower(strings.TrimSpace(tmp[1]))
	if !sl.CheckCmd(in.object.command) {
		in.messageIn = nil
		in.lenIn = 0
		in.MessageOut = "command not found"
		return
	}

	if !sl.CheckKey(in.object.key) {
		in.MessageOut = "key not found"
		return
	}

	in.object.status = true
	in.MessageOut = "работайте, братья"
	in.lenIn++
}

func (in *input) CheckData() {
	if !in.object.status {
		return
	}
	if in.object.key == "вес" {
		in.object.status = sl.CheckNumber(in.object.data)
		if !in.object.status {
			in.MessageOut = fmt.Sprintf("%v - некорректно указан вес\n", in.object.data)
			return
		}
	}

	if in.object.key == "продукт" {
		in.object.status = sl.CheckProd(in.object.data)
		if !in.object.status {
			in.MessageOut = fmt.Sprintf("некорректно указан продукт -%v\n", in.object.data)
			return
		}
	}

	in.MessageOut = "работаем, братья"
}

func main() {
	os.Chdir("../DataBase")
	data, _ := os.ReadFile("tmp.txt")
	sb := strings.Builder{}

	for _, b := range data {
		sb.WriteByte(b)
	}
	var in input
	in.CreateRequest(sb.String())
	in.CheckData()

	fmt.Println(in)

}
