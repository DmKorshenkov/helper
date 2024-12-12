package product

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"unsafe"

	helper "github.com/DmKorshenkov/helper"
	"github.com/DmKorshenkov/helper/bot/food"
	ev "github.com/DmKorshenkov/helper/bot/food"
)

type Product struct {
	Name string
	Food ev.EnergyValue
}

func Prod(message string) {
	message = strings.ToLower(message)
	//var p Product
	//SetProduct(message)
	GetProduct(message)
	//fmt.Println(p.Food)

}

func SetProduct(message string) {
	slice := helper.SplitTrimCheck(message, 0)
	if slice == nil {
		return
	}

	var product Product

	product.Name = strings.ToLower(slice[0])
	product.Food = food.Energy_Value(slice[1:]...)
	setJson(product.Name, product.Food)
}

func GetProduct(name string) *Product {
	err := os.Chdir("./DataBase")
	check(err)
	data, err := os.ReadFile("product.json")
	check(err)
	if len(data) == 0 {
		log.Println("product.json empty")
		return nil
	}

	var tmp = make(map[string]ev.EnergyValue)
	check(json.Unmarshal(data, &tmp))
	var ok bool
	var p = new(Product)

	p.Name = name
	p.Food, ok = search[string, ev.EnergyValue](name, tmp)
	if !ok {
		answer := fmt.Sprintf("%s не найден, возможно такого продукта нет", name)
		fmt.Println(answer)
	}
	return p

}

func search[K comparable, V comparable](name K, maps map[K]V) (V, bool) {
	if res, ok := maps[name]; ok {

		return res, true
	} else {
		return res, false
	}
}

func setJson[K comparable, V any](key K, val V) {
	check(os.Chdir("./DataBase"))
	f, err := os.OpenFile("product.json", os.O_RDWR, 0666)
	check(err)

	data, err := os.ReadFile(f.Name())
	check(err)

	var tmp = make(map[K]V)
	if len(data) != 0 {
		check(json.Unmarshal(data, &tmp))
		data = nil
	}

	tmp[key] = val
	data, err = json.MarshalIndent(tmp, "", "	")
	check(err)
	_, err = f.Write(data)
	check(err)

	check(f.Close())
	fmt.Println("записал product")
}

func check(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}

func size(t any) {
	size := unsafe.Sizeof(t)
	fmt.Printf("size - %d\n", size)
}
