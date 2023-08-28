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
	return fmt.Sprintf("%Q\n", it)
}

func (it Item) Format(f fmt.State, verb rune) {
	val := it.Name
	if verb == 81 {
		val = "Ad : " + val + "\nFiyat :" + fmt.Sprintf("%v", it.Price) + "\nIndirimli Fiyat :" + fmt.Sprintf("%v", it.Price-it.Discount)
	}
	fmt.Fprintf(f, val)
}
