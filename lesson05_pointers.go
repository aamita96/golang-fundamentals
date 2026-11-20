package main

import "fmt"

func pointerCheck() {
	name := "lucia"

	fmt.Println("memory address of name is ", &name)

	m := &name // we can store the memory address of name to another variable.

	fmt.Println("memory address:", m)

	// we can dereference the memory address that we store in the variable using asteric (*)
	fmt.Println("value at memory address:", *m)

	updateName(m)
	fmt.Println("value after update at memory address:", *m)
}

func updateName(x *string) {
	*x = "jason"
}
