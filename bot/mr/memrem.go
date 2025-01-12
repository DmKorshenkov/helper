package mr

import (
	"fmt"

	"github.com/DmKorshenkov/helper/bot/o"
)

type RemMem struct {
}

func (omit *RemMem) Rem(v any) {
	fmt.Printf("v: %v\n", v)
	switch val := v.(type) {
	case o.Weight:
		fmt.Println("Rem type Weight")
		o.RemWeight(val)
	case o.Food:
		fmt.Println("Rem type Food")
		o.RemProd(val)
	default:
		fmt.Println("not type weight/food")
	}
}

func (omit *RemMem) Mem(v any) {
	fmt.Println(v)
}
