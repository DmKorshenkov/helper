package ymd

import "fmt"

func DateFromMap[K comparable, V any](mp map[int]map[int]map[int]map[K][]V, date int) {
	fmt.Println(ConvDateYMD(date))
	if d := date; date > 0 && date <= 31 {
		mp_d := valFromMapDay(mp, d)
		fmt.Println("day == ", d)
		fmt.Println(mp_d)
		return
		//ch<-day
	}
	if m := date / 100; date >= 100 && date <= 1231 {
		mp_m := valFromMapMonth(mp, m)
		fmt.Println("month == ", m)
		fmt.Println(mp_m)
		if d := date % 100; date%100 != 0 && mp_m != nil {
			//ch <- day
			fmt.Println("day == ", d)
			fmt.Println(valFromMapDay(mp, d))
		}
		return
	}
	if y := date / 10000; date%10000 == 0 {
		mp_y := valFromMapYear(mp, y)
		fmt.Println(mp_y)
		return
	}
	y := date / 10000
	mp_y := valFromMapYear(mp, y)
	//	fmt.Println(mp_y)
	if m := date % 10000; m != 0 && mp_y != nil {
		mp_m := valFromMapMonth(mp, m/100)
		fmt.Println(mp_m)
		if d := m % 100; d != 0 && mp_m != nil {
			mp_d := valFromMapDay(mp, d)
			fmt.Println(mp_d)
		}
		return
	}
	fmt.Println("nothing")
	///////////
	///////////
}

func valFromMapDay[K comparable, V any](mp2 map[int]map[int]map[int]map[K][]V, day int) map[K][]V {
	y, m, _ := ConvDateYMD(ConvDateNow())
	if mp2[y] == nil || mp2[y][m] == nil || mp2[y][m][day] == nil {
		fmt.Println("mp[day]==false(not found in map)")
		return nil
	}
	return mp2[y][m][day]
}

func valFromMapMonth[K comparable, V any](mp2 map[int]map[int]map[int]map[K][]V, m int) map[int]map[K][]V {
	y, _, _ := ConvDateYMD(ConvDateNow())
	if mp2[y] == nil || mp2[y][m] == nil {
		fmt.Println("mp[month]==false(not found in map)")
		return nil
	}
	return mp2[y][m]
}

func valFromMapYear[K comparable, V any](mp2 map[int]map[int]map[int]map[K][]V, day int) map[int]map[int]map[K][]V {
	y, _, _ := ConvDateYMD(ConvDateNow())
	if mp2[y] == nil {
		fmt.Println("mp[year]==false(not found in map)")
		return nil
	}
	return mp2[y]
}
