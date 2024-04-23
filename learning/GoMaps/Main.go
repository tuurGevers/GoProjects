package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex
var TestMap = map[string]Vertex{
	"TestEntry": {
		500.0, 200.0,
	},
}

func main() {
	//basic map
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	//map literals
	var mLiteral = map[string]Vertex{
		"Bell Labs": {
			40.68433, -74.39967,
		},
		"Google": {
			37.42202, -122.08408,
		},
	}
	mLiteral["Test"] = Vertex{
		10.0, 34.0,
	}

	fmt.Println(mLiteral["Test"])

	manipulationTest(TestMap)

}
