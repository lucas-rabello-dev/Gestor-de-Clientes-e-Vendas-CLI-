package models

type Product struct {
	ID string
	Name string
	Price uint
}

func (Product) NameTableP() string {
	return "product_table"
}