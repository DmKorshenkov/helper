package main

import (
	"encoding/json"
	"os"

	"github.com/DmKorshenkov/helper/bot/ev"
	"github.com/DmKorshenkov/helper/bot/o"
	"github.com/DmKorshenkov/helper/bot/sl"
)

func main() {
	os.Chdir(".././o")
	//BackRate()
	//memRate(remRate().DiffEv(*ev.SetEv(10, 10, 300, 0)))
	//fmt.Println(rate)

	//	BackRate()
	//_ = rate
}

func setRate(rate o.Object) {
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

func remRate() *ev.EnergyValue {
	data, err := os.ReadFile("ratetmp.json")
	sl.CheckErr(err)
	var rate ev.EnergyValue
	sl.CheckErr(json.Unmarshal(data, &rate))
	return &rate
}

func memRate(rate ev.EnergyValue) {
	f, err := os.OpenFile("ratetmp.json", os.O_WRONLY|os.O_TRUNC, 0666)
	sl.CheckErr(err)

	data, _ := json.MarshalIndent(rate, "", "	")
	f.Write(data)
	f.Close()
}

func BackRate() {
	data, _ := os.ReadFile("rate.json")

	var tmp = o.NewO()
	json.Unmarshal(data, tmp)
	//fmt.Println(tmp)
	data, _ = json.MarshalIndent(tmp.EnergyValue, "", "	")

	f, err := os.OpenFile("ratetmp.json", os.O_WRONLY|os.O_TRUNC, 0666)
	sl.CheckErr(err)
	f.Write(data)
	f.Close()
}
