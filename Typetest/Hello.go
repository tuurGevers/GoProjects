package main

import "fmt"

func main() {
	v := "43" // change me!
	Type := fmt.Sprintf("%T", v)
	switch Type {
	case "string":
		fmt.Printf("string baby")

	default:
		fmt.Print("zeker geen string")
	}
}
