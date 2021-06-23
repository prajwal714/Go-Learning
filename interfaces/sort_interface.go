package main

import (
	"fmt"

	"./mysort"
)

type Person struct {
	firstName, lastName string
}

type Persons []Person

func (p Persons) Len() int { return len(p) }

func (p Persons) Less(i, j int) bool {
	if p[i].firstName == p[j].firstName {
		return p[i].lastName < p[j].lastName
	}
	return p[i].firstName < p[j].firstName
}

func (p Persons) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {

	pArr := Persons{{"Prajwal", "Singh"}, {"Aman", "Rai"}, {"Prajwal", "Singhz"}, {"Abhay", "Chaturvedi"}}
	mysort.Sort(pArr)

	fmt.Println(pArr)

}
