package main

import (
	"log"

	"simplemath.com/utils"
	"um6p.ma/hello/mathutils"
)

func main() {
	a, b := 50, 26
	log.Println("Result is ", mathutils.Add(a, b))
	log.Println("Result is ", utils.Square(a))

}
