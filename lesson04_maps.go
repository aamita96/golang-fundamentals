package main

import "fmt"

// strings as keys and numbers as values
var menu = map[string]float64{
	"soup":           4.99,
	"pie":            7.99,
	"salad":          6.99,
	"toffee pudding": 3.55,
}

// ints as key type
var phonebook = map[int]string{
	1111111111: "mario",
	2222222222: "luigi",
	4444444444: "peach",
	8888888888: "bowser",
}

func print_menu() {
	// fmt.Println(menu)
	// fmt.Println(menu["pie"])

	// looping maps

	for k, v := range menu {
		fmt.Printf("%v - $%v \n", k, v)
	}
}

func open_phonebook() {
	for k, v := range phonebook {
		fmt.Printf("%v - %v\n", k, v)
	}
}
