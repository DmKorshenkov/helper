package o

type Prod struct {
	Name   string
	Weight float64
}

func NewProd() *Prod {

	return &Prod{}
}
func (p *Prod) SetName(name string) Prod {
	p.Name = name
	return *p
}

func (p *Prod) SetWeight(weight float64) Prod {
	p.Weight = weight
	return *p
}

func (p *Prod) SetProd(name string, weight float64) Prod {
	p.Name = name
	p.Weight = weight
	return *p
}
