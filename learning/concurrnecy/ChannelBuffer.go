package main

import "fmt"

func mainBuffer() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	//ch <- 3 overfill
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
