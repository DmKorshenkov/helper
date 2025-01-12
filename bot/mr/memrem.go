package mr

import (
	"fmt"

	"github.com/DmKorshenkov/helper/bot/o"
)

type RemMem struct {
}

func (omit *RemMem) Rem(v any) {
	switch val := v.(type) {
	case o.Weight:
		val.RemWeight()
	case o.Food:
		o.RemProd(val)
	}
}

func (omit *RemMem) Mem(v any) {
	fmt.Println(v)
}
