package main

import "fmt"

var TestMap = map[string]Vertex{
	"TestEntry": {
		500.0, 200.0,
	},
}

func manipulationTest() {
	fmt.Println("map manipulation test")
	printEntryIfExists()
	delete(TestMap, "TestEntry")
	printEntryIfExists()
	TestMap["TestEntry"] = Vertex{
		200.0, 500.0,
	}
	printEntryIfExists()

}

func printEntryIfExists() {
	if elem, ok := TestMap["TestEntry"]; ok {
		fmt.Printf("TestEntry  %v %v \n", ok, elem)
	} else {
		fmt.Printf("TestEntry  %v\n", ok)
	}
}
