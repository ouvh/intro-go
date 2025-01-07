package main

import (
	"fmt"
	"log"

	"github.com/imusmanmalik/randomizer"
	"simplemath.com/utils"
	"um6p.ma/hello/mathutils"
)

func main() {
	a, b := 50, 26
	log.Println("Result is ", mathutils.Add(a, b))
	log.Println("Result is ", utils.Square(a))
	ra, _ := randomizer.RandomInt(100, 555)
	log.Println("Result is ", ra)

	var x = []int{}
	fmt.Println(x == nil)

	var y = map[int]float32{}
	fmt.Println(y)

}
