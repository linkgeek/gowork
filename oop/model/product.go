package model

type Product struct {
	name  string
	price float32
}

func (this *Product) SetName(_name string) {
	this.name = _name
}

func (this *Product) GetName() string {
	return this.name
}

func (this *Product) SetPrice(_price float32) {
	this.price = _price
}

func (this *Product) GetPrice() float32 {
	return this.price
}
