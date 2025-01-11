package memrem

import "fmt"

type RemMem interface {
	Rem(any)
	Mem(any)
}

type empty struct {
}

func (omit *empty) Rem(v any) {
	var empty empty
	_ = empty
	fmt.Println(v)
}

func (omit *empty) Mem(v any) {
	fmt.Println(v)
}
