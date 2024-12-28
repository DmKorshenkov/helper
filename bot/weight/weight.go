package weight

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/DmKorshenkov/helper/bot/sl"
)

type Weight struct {
	//	Time   [2]int  `json:"Time"`
	Weight float64 `json:"Weight"`
	Info   string  `json:"Info,omitempty"`
}

func NewWeight(w float64, inf string) *Weight {
	/*var tmp [2]int
	tmp[0], tmp[1], _ = time.Now().Clock()
	return Weight{Time: tmp, Weight: w, Info: inf}*/
	return &Weight{Weight: w, Info: inf}
}

/*
func (w *Weight) MemWeight()
map[sl.Year][sl.Month][sl.MonthDay]
*/

func (w *Weight) Str() string {
	return fmt.Sprintf("Weight - %v\nInfo - %v\n", w.Weight, w.Info)
}

func (w *Weight) MemWeight(year int, month int, day int) {
	data, err := os.ReadFile("weight.json")
	sl.CheckErr(err)

	if len(data) != 0 {
		mapYear := sl.CmapYear[Weight]()
		sl.CheckErr(json.Unmarshal(data, &mapYear))
		sl.DataInMap(mapYear, *w, year, month, day)

		data, err = json.MarshalIndent(mapYear, "", "	")
		sl.CheckErr(err)
	} else {
		mapYear := sl.CmapYear[Weight]()
		sl.DataInMap(mapYear, *w, year, month, day)
		data, err = json.MarshalIndent(mapYear, "", "	")
		sl.CheckErr(err)
	}

	f, err := os.OpenFile("weight.json", os.O_RDWR, 0666)
	sl.CheckErr(err)

	_, err = f.Write(data)
	sl.CheckErr(err)

	sl.CheckErr(f.Close())
	log.Println("weight set in json compleat")
}

func (w *Weight) RemWeight(year int, month int, day int) []Weight {
	data, err := os.ReadFile("weight.json")
	sl.CheckErr(err)

	var mp = make(map[int]map[int]map[int][]Weight)
	sl.CheckErr(json.Unmarshal(data, &mp))
	slWeight := sl.DataFromMap(mp, year, month, day)
	return slWeight
}

func searchYMD(key int, mp map[int]map[int]map[int][]Weight) (bool, []Weight) {
	_, ok := mp[sl.Year()]
	if !ok {
		return ok, nil
	}
	_, ok = mp[sl.Year()][sl.Month()]
	if !ok {
		return false, nil
	}
	slice, ok := mp[sl.Year()][sl.Month()][key]
	if !ok {
		return ok, nil
	}
	return ok, slice
}
