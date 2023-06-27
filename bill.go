package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// Creates new bill.
func createNewBill(name string) bill {
	billData := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return billData
}

// Pretty prints the bill.
func (b bill) format() string {
	fs := "Bill breakdown: \n\n"
	var total float64 = 0

	for key, val := range b.items {
		fs += fmt.Sprintf("%-20v ...$%v \n", key+":", val)
		total += val
	}

	fs += fmt.Sprintf("\n%-20v ...$%0.2f \n", "tip:", b.tip)
	fs += fmt.Sprintf("\n%-20v ...$%0.2f", "total:", total)

	return fs
}

func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

// Add an item to the bill.
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}
