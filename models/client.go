package models

type Client struct {
	ID string
	Name string
	Email string
	PhoneNumber int
}
// retorna o nome da tabela
func (Client) NameTableC() string {
	return "client_table"
}
