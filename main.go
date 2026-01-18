package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getUserInput(prompt string, r *bufio.Reader) (string, error) {
	// print prompt
	fmt.Print(prompt)
	
	input, err := r.ReadString('\n')		// break on 'enter' key press
	return strings.TrimSpace(input), err 
}

func createNewBill() Bill {
	// read user input
	reader := bufio.NewReader(os.Stdin)		// read from console
	
	name, _ := getUserInput("Create a new bill name: ", reader)
	
	b := generateBill(name)
	fmt.Println("Created a bill - ", b.name)
	
	return b
}

func promptOptions(b Bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getUserInput("Choose option (a - add item, s - save bill, t - add tip, q - quit) : ", reader)
	
	switch opt {
		case "a":
			name, _ := getUserInput("Enter item name: ", reader)
			price, _ := getUserInput("Enter item price: ", reader)
			
			// parse price value to float from string
			p, err := strconv.ParseFloat(price, 64)
			if err != nil {
				fmt.Println("The price must be a number!")
				promptOptions(b)
			}
			
			b.addItem(name, p)
			
			fmt.Println("Item added in the bill.", name, p)
			promptOptions(b)
		case "s":
			b.save()
		case "t":
			tip, _ := getUserInput("Enter tip amount (in $): ", reader)
			
			// parse price value to float from string
			t, err := strconv.ParseFloat(tip, 64)
			if err != nil {
				fmt.Println("The price must be a number!")
				promptOptions(b)
			}
			
			b.addTip(t)
			
			fmt.Println("Tip added to the bill.", t)
			promptOptions(b)
		case "q":
			fmt.Println("You chose to quit the bill system.")
			os.Exit(3)
		default:
			fmt.Println("You can only choose option from: a - add item, s - save bill, t - add tip, q - quit")
			promptOptions(b)		// recyle the asking function
	}
}

func main(){
	myBill := createNewBill()
	promptOptions(myBill)
}