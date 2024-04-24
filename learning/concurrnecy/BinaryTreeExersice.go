package main

import (
	"fmt"
	"sort"
	"sync"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int, quit chan bool) {
	select {
	case <-quit:
		return
	default:
		if t.Left != nil {
			Walk(t.Left, ch, quit)
		}
		if t.Right != nil {
			Walk(t.Right, ch, quit)
		}
		ch <- t.Value

		fmt.Println(t.Value)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	quit := make(chan bool)

	go Walk(t1, ch1, quit)
	go Walk(t2, ch2, quit)

	//store results in array before comparing
	t1c := make([]int, 0, 10)
	t2c := make([]int, 0, 10)

	for i := 0; i < 10; i++ {
		t1c = append(t1c, <-ch1)
		t2c = append(t2c, <-ch2)
	}

	//sort the arrays using sorting lib
	sort.Ints(t1c)
	sort.Ints(t2c)

	//compare the arrays
	for i := 0; i < 10; i++ {
		if t1c[i] != t2c[i] {
			return false
		}
	}

	return true
}

func mainBInary() {
	ch := make(chan int)
	quit := make(chan bool)
	var wg sync.WaitGroup
	wg.Add(1)

	// read values from the channel
	go func() {
		defer wg.Done()
		v := make([]int, 0, 10)
		for i := 0; i < 10; i++ {
			v = append(v, <-ch)
		}
		fmt.Println(v)
	}()

	Walk(tree.New(1), ch, quit)
	close(ch)

	wg.Wait()
}

func main() {
	// mainGoroutines()
	//mainChannels()
	// mainBuffer()
	// mainSelect()
	// mainSelectDefault()
	// mainBInary()
	// fmt.Println(Same(tree.New(1), tree.New(1)))
	// fmt.Println(Same(tree.New(1), tree.New(2)))
	// mainMutex()
	mainCrawler()
}
