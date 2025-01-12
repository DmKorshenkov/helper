package o

import (
	"encoding/json"
	"os"

	"github.com/DmKorshenkov/helper/bot/sl"
)

type Rate struct {
	Name        string
	EnergyValue Ev
}

func NewRate() Rate {
	return Rate{}
}

func (rate *Rate) SetRate(name string, p, f, c, fb, w float64) bool {
	ev := NewEv()
	ev.SetEv(p, f, c, fb)
	_ = ev.SetPortion(w)
	rate.Name = name
	rate.EnergyValue = *ev
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

func RemRate(rate Ev) {
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
	sl.CheckErr(json.Unmarshal(data, &rate))
	return &rate.EnergyValue
}

func BackRate() {
	data, _ := os.ReadFile("rate.json")
	var tmp = NewRate()
	json.Unmarshal(data, &tmp)
	data, _ = json.MarshalIndent(tmp.EnergyValue, "", "	")

	f, _ := os.OpenFile("ratetmp.json", os.O_WRONLY|os.O_TRUNC, 0666)

	_, _ = f.Write(data)
	_ = f.Close()
}
