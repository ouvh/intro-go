package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {

	var low, high int32 = 0, 100
	var bestscore int32 = 1000000000
	var maximalAttempts int32 = 100000000
	fmt.Println("Welcome to Guessing Game")
	fmt.Println("Choose the difficulty:")

	// chosing the difficulty
	for {
		var input string
		var exit int = 0
		fmt.Printf("easy (1-50)\nmedium (1-100)\nhard (1-200)\n")
		fmt.Printf("Enter difficulty(easy , medium , hard):")
		fmt.Scan(&input)

		switch {
		case input == "easy":
			low, high = 1, 50
			exit = 1
		case input == "medium":
			low, high = 1, 100
			exit = 1
		case input == "hard":
			low, high = 1, 200
			exit = 1
		default:
			fmt.Println("Invalid choice try again !!!")

		}
		if exit == 1 {
			break
		}
	}

	// chosing the attempts limit
	for {
		var input string
		fmt.Printf("Enter maximum number of attemps allowed :")
		fmt.Scan(&input)

		var t, err = strconv.ParseInt(input, 10, 32)
		if err != nil {
			fmt.Println("Invalid Number try again !!!")
			continue
		} else {
			maximalAttempts = int32(t)
			break
		}

	}

	for i := 1; ; i++ {
		var chosen int32 = int32(rand.Intn(int(high-low))) + low
		var attempts int32 = 0
		fmt.Printf("Round %d\n", i)
		fmt.Printf("i have chosen a number between %d and %d , let's see if you can guess it: \n", low, high)
		// the game
		for j := 1; j <= int(maximalAttempts); j++ {
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
				bestscore = min(bestscore, attempts)
				fmt.Printf("Number of attempts : %d , best Score : %d", attempts, bestscore)
				break
			}

			if j == int(maximalAttempts) {
				fmt.Printf("You have Lost , Best Score: %d\n", bestscore)
				break
			}

		}
		var continueto int32 = 0
		for {
			var input string
			var exit int = 0
			fmt.Printf("\nContinue (yes or no):")
			fmt.Scan(&input)

			switch {
			case input == "yes":
				continueto = 1
				exit = 1
			case input == "no":
				continueto = 0
				exit = 1
			default:
				fmt.Println("Invalid choice try again !!!")

			}
			if exit == 1 {
				break
			}
		}

		if continueto == 0 {
			fmt.Println("See you Next Time")
			break
		}

	}

}
