package main

import "fmt"

//try printing test entry 3 times
func manipulationTest(testMap map[string]Vertex) {
	//basic fetch will succeed
	fmt.Println("map manipulation test")
	printEntryIfExists(testMap, "TestEntry")

	//entry is deleted so will print failed string
	delete(TestMap, "TestEntry")
	printEntryIfExists(testMap, "TestEntry")

	//put new entry with different values will succeed
	TestMap["TestEntry"] = Vertex{
		200.0, 500.0,
	}
	printEntryIfExists(testMap, "TestEntry")

}

//prints the value of key in testMap if it exists els print fail statement
func printEntryIfExists(testMap map[string]Vertex, key string) {
	if elem, ok := TestMap["TestEntry"]; ok {
		fmt.Printf("%s: %v (Long: %0.1f - Lat: %0.1f) \n", key, ok, elem.Long, elem.Lat)
	} else {
		fmt.Printf("%s:  %v\n", key, ok)
	}
}
