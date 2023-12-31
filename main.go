package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, reader *bufio.Reader) (string, error) {
	fmt.Print(prompt)

	name, err := reader.ReadString('\n')

	return strings.TrimSpace(name), err
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	option, _ := getInput("Choose option [a] add [s] save [t] add tip: ", reader)

	switch option {
	case "a", "A":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)

		fPrice, err := strconv.ParseFloat(price, 63)

		// this is keeping invalid inputs on the stack and it's weird as hell
		if err != nil {
			fmt.Println("Price must be a number...")
			promptOptions(b)
			break
		}

		b.addItem(name, fPrice)

		fmt.Println("Item added...", name, fPrice)
		promptOptions(b)
	case "s", "S":
		fmt.Println("Saving bill...")
		b.save()
	case "t", "T":
		tip, _ := getInput("Tip percentage (don't be a jerk!): ", reader)

		fTip, err := strconv.ParseFloat(tip, 64)

		if err != nil {
			fmt.Println("Tip must be a number...")
			promptOptions(b)
			break
		}

		if fTip < 18.0 {
			fmt.Println("You were a jerk!")
		}

		b.addTip(fTip)
		promptOptions(b)
	default:
		fmt.Println("Invalid option, try again...")
		promptOptions(b)
	}
}

func main() {
	// TODO maybe make reader here so i don't need to create new ones
	myBill := createBill()
	promptOptions(myBill)
}
