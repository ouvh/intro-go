package main

import (
	"fmt"
	"strconv"
)

func main() {
	var sum float32 = 0
	var count uint = 0

	for {

		fmt.Print("Insert a grade or quit (q): ")
		var input string
		var x float32

		fmt.Scan(&input)
		if input == "q" {
			break
		} else {
			var t, err = strconv.ParseFloat(input, 32)
			if err != nil {
				fmt.Println("Invalid Number")
				continue
			} else {
				x = float32(t)
			}
		}
		if x > 100 || x < 0 {
			fmt.Println("Invalid Grade ,Grade should be between 0 and 100")
		} else {
			count++
			sum += x
		}
	}
	var Average float32
	if count != 0 {
		Average = sum / float32(count)
	} else {
		Average = 0
	}
	var letter string
	switch {
	case Average < 60:
		letter = "F"
	case Average < 70:
		letter = "D"
	case Average < 80:
		letter = "C"
	case Average < 90:
		letter = "B"
	case Average >= 90:
		letter = "A"
	}

	fmt.Printf("Average Grade: %f and its letter is %s", Average, letter)

}
