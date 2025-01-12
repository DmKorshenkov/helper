package ymd

import (
	"fmt"
	"sort"
)

func SortmpY[K comparable, V any](mp map[int]map[int]map[int]map[K][]V) (y, m, d int) {
	var keys sort.IntSlice
	//	fmt.Println(mp)
	for k, _ := range mp {
		keys = append(keys, k)

	}

	keys.Sort()
	last := keys.Len() - 1
	//	fmt.Println("last yer in the map - ", keys[last])
	y = keys[last]
	m = sortmpM(mp[y])
	d = sortmpD(mp[y][m])
	return y, m, d
}

func sortmpM[K comparable, V any](mp map[int]map[int]map[K][]V) int {
	var keys sort.IntSlice
	//	fmt.Println(mp)
	for k, val := range mp {
		keys = append(keys, k)
		_ = val
	}

	keys.Sort()
	last := keys.Len() - 1
	//	fmt.Println("last month in the map - ", keys[last])

	sortmpD(mp[keys[last]])
	return keys[last]

}

func sortmpD[K comparable, V any](mp map[int]map[K][]V) int {
	var keys sort.IntSlice
	//	fmt.Println(mp)
	for k, val := range mp {
		keys = append(keys, k)
		_ = val
	}

	keys.Sort()
	last := keys.Len() - 1
	//	fmt.Println("last day in the map - ", keys[last])
	//fmt.Println(mp[last])
	//	sortmpK(mp[keys[last]])
	return keys[last]

}

func sortmpK[K comparable, V any](mp map[K][]V) {

	fmt.Println(mp)
}
