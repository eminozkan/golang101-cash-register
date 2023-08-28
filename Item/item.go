package item

import (
	"fmt"
)

type Item struct {
	Name     string
	Price    float64
	Discount float64
}

func (it Item) Description() string {
	str := "Ad:" + it.Name + ", Fiyat :" + fmt.Sprintf("%v", it.Price) + ", Indirimli Fiyat : " + fmt.Sprintf("%v", it.Price-it.Discount)
	return str
}
