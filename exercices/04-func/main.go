package main

import (
	"fmt"
	"fun/area"
)

func main() {
	var x, y int

	fmt.Scan(&x)
	fmt.Scan(&y)

	a, err := area.CalculateArea(x, y)
	if !err {
		fmt.Println(a)
	} else {
		fmt.Println("Error")
	}
	addingtownumber := add

	var fewewofwefpo func(int, int) int = add
	_ = fewewofwefpo(10, 8)

	fmt.Println(addingtownumber(10, 25))

	temp := saythis("hello")
	for i := 0; i < 10; i++ {
		temp()
	}
}

func add(x, y int) int {
	return x + y
}

func saythis(s string) func() {
	i := 0

	return func() {
		fmt.Printf("%s for the %d time\n", s, i)
		i++
	}

}
