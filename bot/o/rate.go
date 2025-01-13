package o

import (
	"encoding/json"
	"os"
	"strconv"

	"github.com/DmKorshenkov/helper/bot/sl"
)

type Rate struct {
	Name        string
	EnergyValue Ev
}

func NewRate(ev Ev) *Rate {
	var rate = new(Rate)
	rate.Name = strconv.Itoa(int(ev.W.Weight))
	rate.EnergyValue = ev
	return rate
}

func RemRate(rate Rate) bool {
	if ok := rate.setRate(); !ok {
		//check err return true/false
		return false
	} // check err return true/false
	return true
}

func (rate Rate) setRate() bool {
	f_rate, _ := os.OpenFile("rate.json", os.O_CREATE|os.O_RDWR, 0666)
	f_rate_tmp, _ := os.OpenFile("ratetmp.json", os.O_CREATE|os.O_RDWR, 0666)

	data, _ := json.MarshalIndent(rate, "", "	")
	_, _ = f_rate.Write(data)

	data, _ = json.MarshalIndent(rate.EnergyValue, "", "	")
	_, _ = f_rate_tmp.Write(data)

	_ = f_rate_tmp.Close()
	err := f_rate.Close()
	return err != nil
}

func RemRateDay(rate Ev) {
	f, err := os.OpenFile("ratetmp.json", os.O_WRONLY|os.O_TRUNC, 0666)
	sl.CheckErr(err)
	data, _ := json.MarshalIndent(rate, "", "	")
	f.Write(data)
	f.Close()
}

func MemRate() *Ev {
	data, err := os.ReadFile("ratetmp.json")
	sl.CheckErr(err)
	var rate Rate
	sl.CheckErr(json.Unmarshal(data, &rate.EnergyValue))
	return &rate.EnergyValue
}

func BackRate() {
	data, _ := os.ReadFile("rate.json")
	var tmp Rate
	json.Unmarshal(data, &tmp)
	data, _ = json.MarshalIndent(tmp.EnergyValue, "", "	")

	f, _ := os.OpenFile("ratetmp.json", os.O_WRONLY|os.O_TRUNC, 0666)

	_, _ = f.Write(data)
	_ = f.Close()
}
