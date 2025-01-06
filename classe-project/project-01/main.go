package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {

	var low, high int32 = 0, 100
	var chosen int32 = int32(rand.Intn(int(high-low))) + low
	var attempts int32 = 0
	fmt.Println(" Welcome to Guessing Game")
	fmt.Printf("i have chosen a number between %d and %d , let's see if you can guess it \n", low, high)

	for {
		var input string
		var guess int32
		fmt.Printf("Enter your guess:")
		fmt.Scan(&input)

		var t, err = strconv.ParseInt(input, 10, 32)
		if err != nil {
			fmt.Println("Invalid Number try again !!!")
			continue
		} else {
			guess = int32(t)
		}

		attempts++
		if guess > chosen {
			fmt.Println("lower")
		} else if guess < chosen {
			fmt.Println("higher")
		} else {
			fmt.Println("you guessed it !!!")
			fmt.Printf("Number of attempts : %d", attempts)
			break
		}

	}
}
