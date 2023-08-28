package main

import (
	item "app/Item"
	"app/describable"
	"fmt"
)

func calculatePrice(i item.Item) float64 {
	return i.Price - i.Discount
}

func totalPrice(it []item.Item) (float64, error) {
	if it == nil {
		return 0.0, fmt.Errorf("item slice is empty")
	}
	result := 0.0
	for _, val := range it {
		result += val.Price - val.Discount
	}
	return result, nil
}

func printDescription(d describable.Describable) {
	fmt.Println(d.Description())
}

func main() {
	item1 := item.Item{
		Name:     "Elma",
		Price:    5.60,
		Discount: 2.40,
	}

	item2 := item.Item{
		Name:     "Armut",
		Price:    4.60,
		Discount: 2.30,
	}

	item3 := item.Item{
		Name:     "Kiraz",
		Price:    8.90,
		Discount: 2.40,
	}

	slice := []item.Item{item1, item2, item3}

	for _, val := range slice {
		fmt.Println(calculatePrice(val))
		printDescription(val)
	}
	if _, err := totalPrice(slice); err != nil {
		fmt.Println(err)
	}
	tPrice, _ := totalPrice(slice)
	fmt.Println("Total Price : ", tPrice)
}
