package main

import (
	"bufio"
	"fmt"
	"os"
)

type bill struct {
	name       string
	items      map[string]float64
	tipPercent float64
}

// bill{ without a space between the l and { is kind of infuriating, but formatters
// are too convenient
func newBill(name string) bill {
	b := bill{
		name:       name,
		items:      map[string]float64{},
		tipPercent: 0.0,
	}

	return b
}

// using go's deref shortcut
func (b *bill) addTip(percentage float64) {
	b.tipPercent = percentage / 100
}

// not using go's shortcut
func (b *bill) addItem(item string, price float64) {
	(*b).items[item] = price
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin) // reader variable is pointer

	name, _ := getInput("What's your name? ", reader)

	b := newBill(name)
	fmt.Printf("Created bill for %s\n", b.name)

	return b
}

func (b *bill) save() {
	data := []byte(b.format())

	if _, err := os.Stat("./bills/" + b.name + ".txt"); err == nil {
		fmt.Println("Bill exists, overwriting...")
	}

	// no spaces between + grrrrrr
	err := os.WriteFile("./bills/"+b.name+".txt", data, 0644)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Bill for %s created...\n", b.name)
}

func (b *bill) format() string {
	var formatted string = fmt.Sprintf("%s's Bill Breakdown: \n", b.name)
	var total float64 = 0

	for item, price := range b.items {
		formatted += fmt.Sprintf("%-25s ...$%.2f\n", item+":", price)
		total += price
	}

	tip := total * b.tipPercent

	formatted += fmt.Sprintf("%-25s ...$%.2f\n", "Tip:", tip)
	formatted += fmt.Sprintf("%-25s ...$%.2f\n", "total:", total+tip)

	return formatted
}
