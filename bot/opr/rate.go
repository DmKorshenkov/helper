package opr

import (
	"encoding/json"
	"os"

	"github.com/DmKorshenkov/helper/bot/ev"
	"github.com/DmKorshenkov/helper/bot/sl"
)

func SetRate(rate Object) {
	f, err := os.OpenFile("rate.json", os.O_RDWR, 0666)
	sl.CheckErr(err)
	f2, err := os.OpenFile("ratetmp.json", os.O_RDWR, 0666)
	sl.CheckErr(err)
	//	fmt.Println(rate.EnergyValue)
	data, _ := json.MarshalIndent(rate, "", "	")
	f.Write(data)
	//	data = nil
	f.Close()
	data, _ = json.MarshalIndent(rate.EnergyValue, "", "	")
	f2.Write(data)
	f2.Close()
}

func RemRate() *ev.EnergyValue {
	data, err := os.ReadFile("ratetmp.json")
	sl.CheckErr(err)
	var rate ev.EnergyValue
	sl.CheckErr(json.Unmarshal(data, &rate))
	return &rate
}

func MemRate(rate ev.EnergyValue) {
	f, err := os.OpenFile("ratetmp.json", os.O_WRONLY|os.O_TRUNC, 0666)
	sl.CheckErr(err)

	data, _ := json.MarshalIndent(rate, "", "	")
	f.Write(data)
	f.Close()
}

func BackRate() {
	data, _ := os.ReadFile("rate.json")

	var tmp = NewO()
	json.Unmarshal(data, tmp)
	//fmt.Println(tmp)
	data, _ = json.MarshalIndent(tmp.EnergyValue, "", "	")

	f, err := os.OpenFile("ratetmp.json", os.O_WRONLY|os.O_TRUNC, 0666)
	sl.CheckErr(err)
	f.Write(data)
	f.Close()
}
