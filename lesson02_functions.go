package main

import (
	"fmt"
	"math"
)

var fruits = []string{"apple", "banana", "peach", "guava", "mangoes", "tomatoes"}

func sayGreetings(s string) {
	fmt.Printf("Good morning %v \n", s)
}

func sayBye(s string) {
	fmt.Printf("Goodbye %v \n", s)
}

func cycleNames(s []string, f func(string)) {
	for _, value := range s {
		f(value)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}
