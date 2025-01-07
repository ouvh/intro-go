package main

import (
	"encoding/json"
	"log"
	"os"
)

type Person struct {
	Name      string
	Age       int
	Salary    int
	Education string
}

type DAO struct {
	People []Person
}

type Resume struct {
	Average_Age                float64
	Yongest_Person_Age         []string
	Oldest_Person_Age          []string
	Average_salary             float64
	Highest_Salary_Person_Name []string
	Lowest_Salary_Person_Name  []string
	Count_Education_Level      map[string]int
}

func main() {
	// opening the file
	file, err := os.ReadFile("./people.json")

	if err != nil {
		log.Fatal(err)
	}
	// reading the file
	dao := DAO{}
	err = json.Unmarshal(file, &dao)
	if err != nil {
		log.Fatal(err)
	}
	persons := dao.People

	if len(persons) == 0 {
		log.Fatal(" No person found")
	}

	Average_age := 0.0
	Average_salary := 0.0
	min_age := 100000000
	max_age := 0

	var min_person []string
	var max_person []string

	min_salary := 100000000
	max_salary := 0
	var min_salary_person []string
	var max_salary_person []string

	var count = make(map[string]int, 10)

	for _, person := range persons {

		_, exist := count[person.Education]

		if exist {
			count[person.Education]++
		} else {
			count[person.Education] = 1

		}

		Average_age = Average_age + float64(person.Age)
		Average_salary = Average_age + float64(person.Salary)

		// updating the age
		if person.Age < min_age {
			min_age = person.Age
		}
		if person.Age > max_age {
			max_age = person.Age
		}

		// updating the salary
		if person.Salary < min_salary {
			min_salary = person.Salary
		}
		if person.Salary > max_salary {
			max_salary = person.Salary
		}

	}

	for _, person := range persons {

		if person.Age == min_age {
			min_person = append(min_person, person.Name)
		}
		if person.Age == max_age {
			max_person = append(max_person, person.Name)

		}
		if person.Salary == max_salary {
			max_salary_person = append(max_salary_person, person.Name)

		}
		if person.Salary == min_salary {
			min_salary_person = append(min_salary_person, person.Name)
		}

	}

	Average_age = Average_age / float64(len(persons))
	Average_salary = Average_salary / float64(len(persons))

	var result = Resume{Average_age, min_person, max_person, Average_salary, max_salary_person, min_salary_person, count}

	f, _ := os.Create("output.json")
	defer f.Close()
	as_json, errr := json.MarshalIndent(result, "", "\t")
	if errr != nil {
		log.Fatal(errr)
	}
	f.Write(as_json)
	log.Println("Operation ended successfully .File created under the name output.json")

}
