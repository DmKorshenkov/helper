package o

type Food struct {
	Name        string `json:"Name"`
	EnergyValue Ev     `json:"Energy Value"`
}

func NewFood() *Food {
	return &Food{}
}

func SetFood(name string, EnergyValue Ev) *Food {
	return &Food{Name: name, EnergyValue: EnergyValue}
}

func (o *Food) SetFood(name string, EnergyValue Ev) {
	o.Name = name
	o.EnergyValue = EnergyValue
}

func (o *Food) SetName(name string) {
	o.Name = name
}

func (o *Food) SetEnergyValue(EnergyValue Ev) {
	o.EnergyValue = EnergyValue
}

func (o *Food) Food_weight(weight float64) *Food {
	//fmt.Println(o.Ev.W.Weight, "- before")
	o.EnergyValue.W.Weight = weight
	//fmt.Println(o.Ev.W.Weight, "- after")
	return o
}
