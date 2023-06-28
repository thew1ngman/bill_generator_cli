package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Entry point of the app.
func main() {
	newBill := createBill()
	promptOptions(newBill)

	// fmt.Println(newBill.format())
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a new bill name: ", reader)

	b := createNewBill(name)
	fmt.Println("Created the bill - ", b.name)

	return b
}

func getInput(prompt string, reader *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')

	return strings.TrimSpace(input), err
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose an option: [a]add item, [s]save bill, [t]add tip: ", reader)

	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)

		priceFloat, err := strconv.ParseFloat(price, 64)

		if err != nil {
			fmt.Println("The price must be a number.")
			promptOptions(b)
		} else {
			b.addItem(name, priceFloat)
			fmt.Printf("Item added: %v - $%0.2f\n", name, priceFloat)
			promptOptions(b)
		}

	case "s":
		b.save()
		fmt.Println("Saved: ", b)
		fmt.Printf("Bill saved in file: %v\n", b.name+".txt")
	case "t":
		tip, _ := getInput("Enter tip amount: ", reader)

		tipFloat, err := strconv.ParseFloat(tip, 64)

		if err != nil {
			fmt.Println("The price must be a number.")
			promptOptions(b)
		} else {
			b.updateTip(tipFloat)
			fmt.Println("Added tip.")
			promptOptions(b)
		}
	default:
		fmt.Println("Invalid option.")
		promptOptions(b)
	}
}
