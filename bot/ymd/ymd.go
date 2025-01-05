package ymd

import (
	"fmt"
)

func ValFromMap[K comparable, V any](mp map[int]map[int]map[int]map[K][]V, ymd int, k K) []V {
	y, m, d := ConvDateYMD(ymd)
	if val, ok := mp[y][m][d][k]; ok {
		return val
	} else {
		return nil
	}
}

func ValInMap[K comparable, V any](mp map[int]map[int]map[int]map[K][]V, ymd int, k K, val V) map[int]map[int]map[int]map[K][]V {
	if mp == nil {
		mp = make(map[int]map[int]map[int]map[K][]V)
	}
	y, m, d := ConvDateYMD(ymd)
	tmp := mp
	if tmp[y] == nil || tmp[y][m] == nil || tmp[y][m][d] == nil || tmp[y][m][d][k] == nil {
		tmp[y] = Mmp(tmp[y], m, d, k, val)
	} else {
		tmp[y][m][d][k] = append(tmp[y][m][d][k], val)
	}
	mp = tmp
	return mp
}
func Mmp[K comparable, V any](mp map[int]map[int]map[K][]V, m, d int, k K, val V) map[int]map[int]map[K][]V {
	if mp == nil {
		mp = make(map[int]map[int]map[K][]V)
	}
	tmp := mp
	if tmp[m] == nil || tmp[m][d] == nil || tmp[m][d][k] == nil {
		//fmt.Println("nil")
		tmp[m] = dmp(tmp[m], d, k, val)
	} else {
		fmt.Println("!nil")
	}
	mp = tmp
	return tmp
}

func dmp[K comparable, V any](mp map[int]map[K][]V, d int, k K, val V) map[int]map[K][]V {
	if mp == nil {
		mp = make(map[int]map[K][]V)
	}
	tmp := mp
	if tmp[d] == nil || tmp[d][k] == nil {
		//	fmt.Println("nil")
		//var tmp = make(map[int][]int)
		tmp[d] = kmp(tmp[d], k, val)
	} else {
		fmt.Println("!nil")
	}

	mp = tmp
	return mp
}

func kmp[K comparable, V any](mp map[K][]V, key K, val V) map[K][]V {
	if mp == nil {
		mp = make(map[K][]V)
	}
	var tmp = make(map[K][]V)
	tmp = mp
	if tmp[key] == nil {
		//	fmt.Println("nil")
		var tmp2 = append(make([]V, 0, 12), val)
		tmp[key] = tmp2
	} else {
		tmp[key] = append(tmp[key], val)
		fmt.Println("!nil")
	}
	mp = tmp
	return mp
}
