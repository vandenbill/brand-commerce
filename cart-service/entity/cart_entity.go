package entity

type CartEntity struct {
	Products_ID []string
	Total       int
	Quantity    []map[string]int
	User_ID     string
}

type CartPostgreRepo interface {
	Save()
	GetByID()
	Edit()
	Remove()
}
