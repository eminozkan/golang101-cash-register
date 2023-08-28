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

func printDescriptions(d []describable.Describable) error {
	if d == nil {
		return fmt.Errorf("empty descriable slice")
	}
	for _, val := range d {
		fmt.Println(val.Description())
	}
	return nil
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

	fmt.Println("Print Items")
	fmt.Println("____________")
	for _, val := range slice {
		fmt.Printf("%Q\n", val)
		fmt.Println("_____________")
	}

	fmt.Println("Print Description")
	fmt.Println("______________")
	for _, val := range slice {
		printDescription(val)
		fmt.Println("_____________")
	}
	fmt.Println("Print Descriptions")
	fmt.Println("_________________")
	if _, err := describable.ConvertItemSliceToDescribableSlice(slice); err != nil {
		fmt.Println(err)
	}
	describableSlice, _ := describable.ConvertItemSliceToDescribableSlice(slice)

	err := printDescriptions(describableSlice)
	if err != nil {
		return
	}

	if _, err := totalPrice(slice); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Total Price")
	fmt.Println("______________")
	tPrice, _ := totalPrice(slice)
	fmt.Println("Total Price : ", tPrice)
}
