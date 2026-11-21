package main

import (
	"fmt"
	"os"
)

/*
*

  - Go (Golang) does not support classes in the traditional OOP sense like Java, C++, or Python.

  - However, Go supports most object-oriented programming concepts using other features.

  - Struct is a blueprint which describes type of data. Its similar to the classes like other programming languages.

  - A struct can be useful for grouping data together to create records.

    OOP Concepts in Traditional OOP Languages vs Go:

    1. Class              -> Struct
    2. Object             -> Struct Instance
    3. Methods             -> Methods with Receivers
    4. Inheritance         -> Composition (Struct Embedding)
    5. Interface           -> Interface (Implicit Implementation)
    6. Constructor         -> Constructor Function
    7. Encapsulation       -> Exported and Unexported Fields/Methods
    8. Polymorphism        -> Interfaces
    9. Abstraction         -> Interfaces
    10. Method Overloading -> Not Supported Directly
    11. Method Overriding  -> Not Supported Directly (Method Shadowing via Embedding)
    12. this / self        -> Receiver (e.g., func (u User) ...)
    13. Static Members     -> Package-level Variables and Functions
    14. Multiple Inheritance -> Not Supported (Composition Instead)
*/
type Bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bill
func newBill(name string) Bill {
	b := Bill{
		name:  name,
		items: map[string]float64{}, // "coffee": 5.99, "sparkling water": 2.99,
		tip:   0.0,
	}

	return b
}

// format the bill
// By adding (b Bill) in front of the function name we're limiting this function to only to bill objects.
func (b *Bill) format() string {

	formattedStr := b.name + " breakdown \n"
	var total float64 = 0
	for k, v := range b.items {
		// -30 adds 30 characters to the right side, doesn't matter how many characters are there.
		// +30 or 30 adds 30 characters to the left side in the same way.
		formattedStr += fmt.Sprintf("%-30v $%v \n", k+":", v)
		total += v
	}

	formattedStr += fmt.Sprintf("%-30v $%v \n", "Tip:", b.tip)
	total += b.tip

	formattedStr += fmt.Sprintf("%-30v $%v \n", "Total:", total)
	fmt.Println(formattedStr)

	return formattedStr
}

// update tip
func (b *Bill) updateTip(tip float64) {
	b.tip = tip
}

func (b *Bill) addItems(key string, value float64) {
	b.items[key] = value
}

// save bill
func (b *Bill) save() {
	data := []byte(b.format())

	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)

	if err != nil {
		panic(err)
	}

	fmt.Println("bill was saved to file")
}

func generateBill() {
	myBill := newBill("lucia")

	myBill.addItems("coffee", 5.66)
	myBill.addItems("cake", 9.99)

	myBill.updateTip(1)

	myBill.format()
	// fmt.Println(myBill)

}
