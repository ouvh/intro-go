package main

import (
	"fmt"

	"example.com/exercice-2/area"
)

func main() {
	var height, width float64
	fmt.Scan(&height)
	fmt.Scan(&width)
	fmt.Printf("Area in meters = %f \n", area.CalculateArea(float64(height), float64(width)))
	fmt.Printf("Area in foot = %f", height*width*3.28*3.28)

	if height == 0 {
		fmt.Print("there is something wrong with the height")
	} else {
		height = 1
	}
	grade := 50

	switch grade {
	case 10:
		fmt.Print("A")
		fallthrough
	case 50:
		fmt.Print("55")
	}

}
