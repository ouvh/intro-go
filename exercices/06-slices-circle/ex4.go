package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Person struct {
	Name string
}
type Employee struct {
	Person
	id int
}

func main() {
	var g string = "oussama"
	fmt.Println(wordFreq(g))

	var h = Person{"rreferg"}
	fmt.Println(h)

	file, err := os.ReadFile("./person.json")

	if err != nil {
		log.Fatal(err)
	}
	p := Person{}
	err = json.Unmarshal(file, &p)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(p)

}

func wordFreq(s string) map[string]uint {
	var fre = make(map[string]uint, 10)
	var tokens []string

	i, j := 0, 0
	for i < len(s) && j < len(s) {
		if s[j] == ' ' {

			i++
			j++
		} else {
			i = j
			for j < len(s) && s[j] != ' ' {
				j++
			}
			tokens = append(tokens, s[i:j])
			i = j + 1
			j++

		}
	}

	for _, token := range tokens {
		_, exist := fre[token]

		if exist {
			fre[token]++
		} else {
			fre[token] = 1
		}
	}

	return fre
}
