package main

import "fmt"

func main(){
	myBill := generateBill("Test Bill")
	
	// add items to the bill
	myBill.addItem("Cheesecake", 1.99)
	myBill.addItem("Fries", 2.39)
	myBill.addItem("Coke", 1.25)
	
	// add tip to the bill
	myBill.addTip(5)
	
	// format and return the formatted bill
	formatedBill := myBill.format()
	fmt.Println(formatedBill)
}