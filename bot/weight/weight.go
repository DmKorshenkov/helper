package weight

import (
	"fmt"
)

type Weight struct {
	Info   string  `json:",omitempty"`
	Weight float64 `json:",omitempty"`
}

func NewW() *Weight {
	return &Weight{}
}

func SetW(w float64, i string) *Weight {
	return &Weight{Weight: w, Info: i}
}

func (w *Weight) SetWeight(weight float64) {
	w.Weight = weight
}

func (w *Weight) SetInfo(info string) {
	w.Info = info
}

func (w *Weight) Str() string {
	if w.Info == "" {
		return fmt.Sprintf("weight - %v", w.Weight)
	}
	return fmt.Sprintf("weight - %v\ninfo - %v", w.Weight, w.Info)
}
