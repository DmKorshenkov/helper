package helper

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func SplitTrimSpaceCheck(get string, n int) []string {
	slice := strings.Split(get, " ")
	if len(slice) == n || n == 0 {
		for in, str := range slice {
			slice[in] = strings.TrimSpace(str)
		}
		return slice
	}
	return nil
}

func SplitPlusCheck(get ...string) {

}

func Split(get string) (slice []string) {
	if get == "" {
		return nil
	}
	slice = strings.Split(get, " ")
	return slice
}

func SplitTrimSpace(get string) (slice []string) {
	if get == "" {
		return nil
	}

	slice = strings.Split(get, " ")
	for in, str := range slice {
		slice[in] = strings.TrimSpace(str)
	}
	return slice
}

func CheckErr(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}

func Type(t any) {
	fmt.Printf("%#T\n%#v\n", t, t)
}

func SetMapInJson[K comparable, V comparable](key K, val V, fileName string) {
	f, err := os.OpenFile(fileName, os.O_RDWR, 0666)
	CheckErr(err)

	data, err := os.ReadFile(f.Name())
	CheckErr(err)

	var tmp = make(map[K]V)

	if len(data) != 0 {
		err = json.Unmarshal(data, &tmp)
		CheckErr(err)
		data = nil
	}
	if SearchKeyInMap(key, &tmp) {
		fmt.Println("Value по такому ключу уже есть в map")
		return
	}
	tmp[key] = val

	_, err = json.MarshalIndent(tmp, "", "	")
	CheckErr(err)

	err = f.Close()
	CheckErr(err)

	fmt.Println("compleat")
}
func GetFromJson(name string) (t any) {

	data, err := os.ReadFile(name)
	CheckErr(err)
	if len(data) == 0 {
		log.Println("func - GetMapFromJson - file empty!")
		return nil
	}

	CheckErr(json.Unmarshal(data, &t))

	return t
}

func SearchKeyInMap[K comparable, V comparable](key K, dic *map[K]V) bool {
	if _, ok := (*dic)[key]; ok {
		return true
	}
	return false
}

func CheckNumber(slice ...string) bool {
	for _, str := range slice {
		for _, r := range str {
			if !unicode.IsDigit(r) && r != '.' {
				return false
			}
		}
	}
	return true
}
