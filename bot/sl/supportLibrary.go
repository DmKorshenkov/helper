package sl

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func CreateTimeData() (t time.Time) {
	return time.Now()
}

func WeekDay() int {
	return int(time.Now().Weekday())
}

func Month() int {
	return int(time.Now().Month())
}

func MonthDay() int {
	return time.Now().Day()
}

func Year() int {
	return int(time.Now().Year()) % 2000
}

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

func SplitSlashN(get string) []string {
	if get == "" {
		return nil
	}
	slice := strings.Split(get, "\n")
	return slice
}

func SplitSpaceTrim(get string) (slice []string) {
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

func SetValInJson[K comparable, V comparable](Key K, Val V, fileName string) {
	f, err := os.OpenFile(fileName, os.O_RDWR, 0666)
	CheckErr(err)

	data, err := os.ReadFile(f.Name())
	CheckErr(err)

	var tmp = make(map[K]V)

	if len(data) != 0 {
		fmt.Println("data != nil")
		err = json.Unmarshal(data, &tmp)
		CheckErr(err)
		data = nil
	}
	/*if SearchKeyInMap(Key, tmp) {
		fmt.Println("Value по такому ключу уже есть в map")
		return
	}*/
	tmp[Key] = Val

	data, err = json.MarshalIndent(tmp, "", "	")
	CheckErr(err)

	_, err = f.Write(data)
	CheckErr(err)

	err = f.Close()
	CheckErr(err)

	fmt.Println("compleat")
}

func GetMapFromJson[K comparable, V comparable](fileName string) map[K]V {
	data, err := os.ReadFile(fileName)
	CheckErr(err)

	if len(data) == 0 {
		log.Println("len data = 0, file empty")
		return nil
	}

	var tmp = make(map[K]V)
	json.Unmarshal(data, &tmp)
	return tmp
}

func GetValInJson[K comparable, V comparable](Key K, filename string) (V, bool) {
	//var tmp = make(map[K]V)
	tmp := GetMapFromJson[K, V](filename)
	return tmp[Key], SearchKeyInMap(Key, tmp)

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

func SearchKeyInMap[K comparable, V comparable](key K, dic map[K]V) bool {
	if _, ok := dic[key]; ok {
		return true
	}
	return false
}

func CheckNumber(str string) bool {
	for _, r := range str {
		if !unicode.IsDigit(r) && r != '.' {
			return false
		}
	}
	return true
}

func ParF(get string) (f64 float64) {
	f64, err := strconv.ParseFloat(strings.TrimSpace(get), 64)
	if err != nil {
		log.Println(err.Error())
	}
	//	f64 = f64 * 100 / 100
	f64 = math.Round(f64*100) / 100
	return f64
}

func CheckCmd(command string) bool {
	if command != "запомни" && command != "вспомни" {
		return false
	}
	return true
}

func CheckKey(key string) bool {
	if key != "вес" && key != "продукт" && key != "прием пищи" {
		return false
	}
	return true
}

func CheckProd(data string) bool {
	var slice = make([]string, 0, 2)
	if len(slice) != 2 {
		return false
	}

	if !CheckNumber(slice[1]) {
		return false
	}

	return true
}
