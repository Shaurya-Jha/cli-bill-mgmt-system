package main

import (
	"fmt"
	"os"
)

type Bill struct {
	name string
	items map[string]float64
	tip float64
}

// create a new bill with bill name
func generateBill(name string) Bill {
	b := Bill{
		name: name,
		items: map[string]float64{},
		tip: 0,
	}
	
	return b
}

// receiver function to add tip to the bill
func (b *Bill) addTip(tip float64) {
	b.tip = tip
}

// receiver function to add an item to the bill
func (b *Bill) addItem(name string, price float64) {
	b.items[name] = price
}

// format and then return the bill output to terminal
func (b Bill) format() string {
	fs := "Bill breakdown\n\n"
	var total float64
	
	// loop map items
	for key, value := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v\n", key+":", value)
		total += value
	}
	
	// tip
	fs += fmt.Sprintf("%-25v ...$%v\n\n", "tip: ", b.tip)
	// total
	fs += fmt.Sprintf("%-25v ...$%0.2f\n", "Total :", total)
	
	return fs
}

// save bill in text file
func (b *Bill) save(){
	// convert bill data into byte size
	data := []byte(b.format())
	
	// save files to the bills folder as per the bill name
	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Bill saved successfully!")
}