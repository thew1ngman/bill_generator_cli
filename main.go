package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	newBill := createBill()
	promptOptions(newBill)

	fmt.Println(newBill)
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a new bill name: ", reader)

	b := createNewBill(name)
	fmt.Println("Created the bill - ", b.name)

	return b
}

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose an option: [a]add item, [s]save bill, [t]add tip: ", reader)

	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)

		conv, _ := strconv.ParseFloat(price, 64)

		b.addItem(name, conv)

		fmt.Printf("%v - $%0.2f\n", name, conv)
	case "s":
		fmt.Printf("You chose to [%v]save the bill.\n", opt)
	case "t":
		tip, _ := getInput("Tip amount: ", reader)
		tipFloat, _ := strconv.ParseFloat(tip, 64)

		b.updateTip(tipFloat)

		fmt.Printf("You chose to [%v]add tip.\n", opt)
	default:
		fmt.Println("That was not in the options.")
		promptOptions(b)
	}
}
