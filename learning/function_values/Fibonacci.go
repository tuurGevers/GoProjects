package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	n1 := 0
	n2 := 0
	num := 0

	return func() int {
		num = n1 + n2
		if num == 0 {
			n2 = 1
		}
		n1 = n2
		n2 = num
		return num
	}
}

func fibonacciMain() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
