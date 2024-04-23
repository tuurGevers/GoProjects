package main

import "fmt"

//should return a new function that added the value to the old funtion
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func counter() (func() int, func(n int) int, func() int, func() int) {
	sum := 0

	getCount := func() int {
		return sum
	}

	setCount := func(amount int) int {
		sum = amount
		return sum
	}

	increment := func() int {
		sum++
		return sum
	}

	decrement := func() int {
		sum--
		return sum
	}

	return getCount, setCount, increment, decrement

}

func mainCounter() {
	getCount, setCount, increment, decrement := counter()
	fmt.Println(getCount())  // 0
	fmt.Println(increment()) // 1
	fmt.Println(increment()) // 2
	setCount(10)
	fmt.Println(getCount())  // 10
	fmt.Println(decrement()) // 10

}

//creates to adder functions as closures and add/subtracts values from them
func mainAdder() {
	//bound the adder function to pos and neg
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
