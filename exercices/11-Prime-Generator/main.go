package main

import "fmt"

func primeGenerator(n int, ch chan<- int) {
	ch <- n
	close(ch)
}

func printPrimes(ch <-chan int) {
	for i := range ch {
		fmt.Println(i)
	}
}
func main() {
	ch := make(chan int)
	go primeGenerator(10, ch)

}
