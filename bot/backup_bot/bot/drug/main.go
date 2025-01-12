package main

type Drug struct {
	Name  string
	Price uint
	Count float64
	Dose  float64
}

type Med struct {
	Dose float64
}

func CreateDrug() *Drug {
	return &Drug{}
}

func (d *Drug) NewDrug(name string, price uint, count float64, dose float64) {
	d.Name = name
	d.Price = price
	d.Count = count
	d.Dose = dose
}

func NewMedic(name string, price uint, count float64, dose float64) *Drug {
	return &Drug{Name: name, Price: price, Count: count, Dose: dose}
}

func main() {
	d := CreateDrug()
	d.NewDrug("azaleptin", 900, 50, 25)

}
