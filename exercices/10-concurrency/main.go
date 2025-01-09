package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func Square(x int) {
	fmt.Println("Square of ", x, " ", x*x)
	defer wg.Done()
}

func main() {
	var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, number := range numbers {
		wg.Add(1)
		go Square((number))
	}

	wg.Wait()

}
