package product

type Product struct {
	Number      int
	Description string
}

func NewProduct(n int, d string) *Product {
	return &Product{
		Number:      n,
		Description: d,
	}
}

func (p *Product) Clone() Cloner {
	c := *p
	return &c
}

func (p *Product) GetDescription() string {
	return p.Description
}

func (p *Product) SetDescription(d string) {
	p.Description = d
}
