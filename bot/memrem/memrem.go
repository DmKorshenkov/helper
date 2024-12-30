package memrem

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/DmKorshenkov/helper/bot/ev"
	"github.com/DmKorshenkov/helper/bot/sl"
	w "github.com/DmKorshenkov/helper/bot/weight"
)

type MR interface {
	Mem(any)
	Rem(any, int)
}

type object struct{}

func (o *object) Mem(input any) {
	switch val := input.(type) {
	case w.Weight:
		f, err := os.OpenFile("weight.json", os.O_RDWR, 0666)
		sl.CheckErr(err)
		data, err := os.ReadFile(f.Name())
		sl.CheckErr(err)

		mp := sl.CmapYear[w.Weight]()
		if len(data) != 0 {
			sl.CheckErr(json.Unmarshal(data, &mp))
		}
		sl.DataInMap(mp, val, sl.Year(), sl.Month(), sl.MonthDay())
		data, err = json.MarshalIndent(mp, "", "	")
		sl.CheckErr(err)

		n, err := f.Write(data)
		sl.CheckErr(err)
		sl.CheckErr(f.Close())
		if err == nil && n == len(data) {
			fmt.Println("true")
		}
	case ev.EnergyValue:
		f, err := os.OpenFile("ev.json", os.O_RDWR, 0666)
		sl.CheckErr(err)
		data, err := os.ReadFile(f.Name())
		sl.CheckErr(err)

		mp := sl.CmapYear[ev.EnergyValue]()
		if len(data) != 0 {
			sl.CheckErr(json.Unmarshal(data, &mp))
		}
		sl.DataInMap(mp, val, sl.Year(), sl.Month(), sl.MonthDay())
		data, err = json.MarshalIndent(mp, "", "	")
		sl.CheckErr(err)

		n, err := f.Write(data)
		sl.CheckErr(err)
		sl.CheckErr(f.Close())
		if err == nil && n == len(data) {
			fmt.Println("true")
		}
	default:
		log.Println("error type")
	}
}

func (o *object) Rem(t any, d int) {
	switch t.(type) {
	case w.Weight:
		data, err := os.ReadFile("weight.json")
		sl.CheckErr(err)

		mp := sl.CmapYear[w.Weight]()
		sl.CheckErr(json.Unmarshal(data, &mp))

		y, m, d := sl.ConvData2(d)
		wSl := sl.DataFromMap(mp, y, m, d)
		for _, w := range wSl {
			fmt.Println(w.Str())
		}
	case ev.EnergyValue:
		data, err := os.ReadFile("ev.json")
		sl.CheckErr(err)

		mp := sl.CmapYear[ev.EnergyValue]()
		sl.CheckErr(json.Unmarshal(data, &mp))

		y, m, d := sl.ConvData2(d)
		eSl := sl.DataFromMap(mp, y, m, d)
		for _, e := range eSl {
			fmt.Println(e.Str())
		}
	default:
		print("error type")
	}
}
