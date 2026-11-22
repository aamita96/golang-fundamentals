package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("I am in main!")
	hello_world()
}

// method
func hello_world() {
	fmt.Println("hello world!")
	fmt.Println("-------------------------------- String --------------------------------")
	variables_examples()
	fmt.Println("-------------------------------- Print & Formatting --------------------------------")
	printing_and_formatting()
	fmt.Println("-------------------------------- Array & Slices --------------------------------")
	arrays_and_slices()
	fmt.Println("-------------------------------- String Package Methods --------------------------------")
	string_package()
	fmt.Println("-------------------------------- Sort Package Methods --------------------------------")
	sort_package()
	fmt.Println("-------------------------------- Loops in Go --------------------------------")
	loops_in_go()
	fmt.Println("-------------------------------- Conditional Statements in Go --------------------------------")
	boolean_and_conditions()
	fmt.Println("-------------------------------- Method calls in Go --------------------------------")
	sayGreetings("Mario")
	sayGreetings("Peach")
	sayBye("Bowser")
	cycleNames(fruits, sayGreetings)
	fmt.Printf("Area of the circle is %0.2f \n", circleArea(30))
	first1, second1 := getInitials("the brown fox ran over the lazy dog!")
	fmt.Println("Initials of the string ", first1, second1)
	first2, second2 := getInitials("lazy frog!")
	fmt.Println("Initials of the string ", first2, second2)
	first3, second3 := getInitials("superman")
	fmt.Println("Initials of the string ", first3, second3)
	fmt.Println("-------------------------------- Maps in Go --------------------------------")
	print_menu()
	phonebook[8888888888] = "yoshi" // This will update the value in the map.
	open_phonebook()
	fmt.Println("-------------------------------- Pointers in Go --------------------------------")
	pointerCheck()
	fmt.Println("-------------------------------- Structs in Go --------------------------------")
	generateBill()
	fmt.Println("-------------------------------- Accepting User Inputs in Go --------------------------------")
	billGenerator()
	fmt.Println("-------------------------------- Interfaces in Go --------------------------------")
	generateShapes()
}

func variables_examples() {
	// Strings
	var name string = "mario"     // explicity typing the variable
	var email = "luigi@gmail.com" // automatically infer type or dynamic type
	var city string
	password := "Mario@1234" // := It is used to declare and initialize a variable in one step.

	fmt.Println("\nIntroduction to variables in Go --")
	fmt.Println("Name: "+name, "Email: "+email, "City: "+city)

	fmt.Println("City: "+city, "Password: "+password)

	// int
	var age int = 55
	var num = 777
	score := 201

	fmt.Println("Age:" + strconv.Itoa(age)) // as number can't be concatinated with string or
	fmt.Println("Age:", age)

	fmt.Println("Number:", num)
	fmt.Println("Score:", score)

	// bits & memory (https://go.dev/ref/spec#Numeric_types)
	var count int16 = 333
	fmt.Println(count)

	// float
	var temp float32 = 33.79
	var avg float64 = 69988888888888.9999
	median := 556677.88990

	fmt.Println(temp)
	fmt.Println(avg)
	fmt.Println(median)
}

func printing_and_formatting() {
	age := 15
	name := "mario"

	// Println
	fmt.Println("My name is ", name, " and my age is ", age)

	// Printf (formatting strings)
	fmt.Printf("My name is %v and my age is %v with Formatted String \n", name, age)
	fmt.Printf("My Age is %v and my name is %q \n", age, name) // String value with quotes
	fmt.Printf("Age variable is of type %T \n", age)

	//Sprintf (save formatted strings)
	var str = fmt.Sprintf("My name is %v \n", name)
	fmt.Println(str)
}

func arrays_and_slices() {
	// var ages [3]int = [3]int{20, 25, 30}
	// 			OR
	var ages = [3]int{20, 25, 30}
	names := [4]string{"Yoshi", "Mario", "Peach", "Bowser"}

	names[2] = "Luigi"

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	/**
	 *	Slices are the arrays like other languages where we can change the length,
	 *	we can add items, take items away from it.
	 *
	 * 	slices use arrays under the hood
	 */

	var scores = []int{10, 20, 30}
	scores = append(scores, 45)

	fmt.Println(scores)

	// slice ranges
	rangeOne := names[1:3]
	fmt.Println(rangeOne)
}

func string_package() {
	greeting := "hello there friends!"
	sentence := "That that is is that that is not is not."
	palin := "radar"

	fmt.Println("Does contain 'fri'? ", strings.Contains(greeting, "fri"))
	fmt.Println("Character count is :", strings.Count(sentence, "is"))

	fmt.Println("Replacing word hello with hi: ", strings.ReplaceAll(greeting, "hello", "hi"))
	fmt.Println("UpperCase String :", strings.ToUpper(greeting))
	// fmt.Println("Titlecase String :", strings.ToTitle(greeting))
	fmt.Println("Index of 'th' :", strings.Index(greeting, "th"))

	fmt.Println(strings.Join(strings.Split(greeting, " "), ","))
	fmt.Println(strings.Compare(palin, ""))

	// the original value is unchanged
	fmt.Println("original string value = ", greeting)
}

func sort_package() {
	ages := []int{20, 25, 15, 78, 90, 48, 30, 25, 29}
	fmt.Println(ages)
	sort.Ints(ages)
	fmt.Println(ages)

	fmt.Println("Are Ages Sorted? ", sort.IntsAreSorted(ages))
	fmt.Println(sort.SearchInts(ages, 30))

	// sort.SearchInts() requires the slice to be sorted in ascending order before searching.
	// If a number doesn't exits in the array/slice then It returns the index where the value would be inserted to maintain sorted order.
	index := sort.SearchInts(ages, 66)
	fmt.Println("66 value doesn't exists ", index)

	names := []string{"Yoshi", "Mario", "Peach", "Bowser", "Luigi"}

	sort.Strings(names)
	fmt.Println(names)
	fmt.Println(sort.SearchStrings(names, "Peach"))
}

func loops_in_go() {
	x := 0

	// While loop equivalent in Go
	for x < 5 {
		fmt.Println("*")
		x++
	}

	// Normal for loop
	for i := 1; i < 10; i++ {
		for j := 1; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	// Do-while loop equivalent
	for {
		fmt.Println("This executes at least once")

		if true {
			break
		}
	}

	names := []string{"yoshi", "mario", "luigi", "peach", "bowser"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// 			OR (similar to for in loop in javascript)

	// for index, value := range names {
	// 	fmt.Printf("%v at index %v \n", value, index)
	// }

	for _, value := range names {
		fmt.Printf("%v at index \n", value)
	}
}

func boolean_and_conditions() {
	age := 45

	fmt.Println(age <= 50) // true
	fmt.Println(age < 40)  // false
	fmt.Println(age == 45) // true
	fmt.Println(age != 50) // true

	if age < 30 {
		fmt.Println("Age is less than 30")
	} else if age < 40 {
		fmt.Println("Age is less than 40")
	} else {
		fmt.Println("Age is not less than 45")
	}

	names := []string{"mario", "luigi", "peach", "bowser"}

	for index, value := range names {
		if index == 1 {
			fmt.Println("continuing at pos", index)
			continue
		}

		if index > 2 {
			fmt.Println("breaking at pos", index)
			break
		}

		fmt.Printf("%v ", value)
	}
	fmt.Println()
}
