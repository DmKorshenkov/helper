package o

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/DmKorshenkov/helper/bot/ymd"
)

type Weight struct {
	Info   string  `json:",omitempty"`
	Weight float64 `json:",omitempty"`
}

func NewW() *Weight {
	return &Weight{}
}

func SetW(weight float64, info string) *Weight {
	return &Weight{Weight: weight, Info: info}
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

func RemWeight(w Weight) {
	f, err := os.OpenFile("weight.json", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err.Error())
	}
	mp := make(map[int]map[int]map[int]map[string][]Weight)
	data, _ := os.ReadFile("weight.json")
	if len(data) != 0 {
		_ = json.Unmarshal(data, &mp)
		fmt.Println(mp)
	}
	_ = ymd.ValInMap[string, Weight](mp, ymd.ConvDateNow(), w.Info, w)
	fmt.Println(mp)
	data, err = json.MarshalIndent(mp, "", "	")
	if err != nil {
		fmt.Println("!")
	}
	n, _ := f.Write(data)
	if n != len(data) {
		fmt.Println("n!=data")
	}
	_ = f.Close()
}
