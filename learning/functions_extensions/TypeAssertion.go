package main

import "fmt"

func mainTypes() {
	do(21)
	do("hello")
	do(true)
}

// do is a function that performs different actions based on the type of the input.
// It takes an interface{} parameter and uses a type switch to determine the type of the input.
// If the input is of type int, it prints the value multiplied by 2.
// If the input is of type string, it prints the length of the string.
// For any other type, it prints "I don't know about type <type>!".
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}
