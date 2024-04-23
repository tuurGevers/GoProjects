package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) (PixelArray [][]uint8) {
	// allocate space for a array of arrayès with length dy
	PixelArray = make([][]uint8, dy)

	//initialize array for each entry in PixelArray
	for i := range PixelArray {
		PixelArray[i] = make([]uint8, dx)
	}

	//fill the array
	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			// Example function: XOR of the indices
			PixelArray[i][j] = uint8(i ^ j)
		}
	}

	return
}

func main() {
	pic.Show(Pic)
}
