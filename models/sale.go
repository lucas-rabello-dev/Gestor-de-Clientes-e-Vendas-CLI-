package models

type Sale struct {
	ID string `gorm: primaryKey`
	Product *Product
	Client *Client
	Date string
	Amount uint
}

func (Sale) NameTableS() string {
	return "sale_table"
}
