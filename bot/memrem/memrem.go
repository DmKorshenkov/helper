package memrem

import (
	"github.com/DmKorshenkov/helper/bot/opr"
	"github.com/DmKorshenkov/helper/bot/weight"
)

type MR interface {
	Mem(any)
	Rem(any)
}

type O struct{}

func (o *O)Rem(vr any) {
	switch val := vr.(type){
	case weight.Weight:
			weight.RemWeight(val)
	case opr.Prod:
		val.
	}
}
