package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Generic input function
func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill() Bill {
	reader := bufio.NewReader(os.Stdin)
	name, _ := getInput("Name: ", reader)

	b := newBill(name)

	fmt.Println("Created a bill - ", b.name)
	return b
}

func promptOptions(b Bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)

	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)

		priceInt, err := strconv.ParseFloat(price, 64)

		if err != nil {
			fmt.Println("Conversion error :", err)
			promptOptions(b)
		}

		b.addItems(name, priceInt)

		fmt.Println("Item added: ", name, priceInt)
		promptOptions(b)

	case "t":
		tip, _ := getInput("Enter tip amount ($): ", reader)

		t, err := strconv.ParseFloat(tip, 64)

		if err != nil {
			fmt.Println("Conversion error :", err)
			promptOptions(b)
		}

		b.updateTip(t)
		fmt.Println("Tip added: ", t)
		promptOptions(b)

	case "s":
		fmt.Println("Generating your bill ....")
		b.save()
	default:
		fmt.Println("invalid input")
		promptOptions(b)
	}
}

func billGenerator() {
	mybill := createBill()

	promptOptions(mybill)

}
