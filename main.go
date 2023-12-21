package main

import "fmt"

func main() {
	myBill := newBill("Bob Fossil")
	myBill.addTip(0.19)
	myBill.addItem("Oysters", 8.99)
	myBill.addItem("Skrimps", 6.99)

	fmt.Println(myBill.format())
}
