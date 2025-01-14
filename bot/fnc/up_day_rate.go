package fnc

import (
	"encoding/json"
	"os"
	"time"

	"github.com/DmKorshenkov/helper/bot/o"
)

func UpDayRate(ch chan string) {
	for {
		var t = time.Now()
		if h, m, s := t.Clock(); h == 0 && m == 0 && s == 0 {
			check := BackRate()
			if check != "" {
				ch <- check
			}
			time.Sleep(time.Hour*23 + time.Minute*59)
		}
	}
}

func BackRate() string {
	data, _ := os.ReadFile("rate.json")
	if len(data) == 0 {
		return "rate not found or BackRate has another problem"
	}
	var tmp o.Rate
	json.Unmarshal(data, &tmp)
	data, _ = json.MarshalIndent(tmp.EnergyValue, "", "	")

	f, _ := os.OpenFile("ratetmp.json", os.O_WRONLY|os.O_TRUNC, 0666)

	_, _ = f.Write(data)
	_ = f.Close()
	return ""
}
