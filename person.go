package main

import "fmt"

type Person struct {
	name      string
	address   string
	age       int
	isStudent bool
}

func newPerson(n string, add string, a int, student bool) *Person {
	temp := Person{name: n, address: add, age: a, isStudent: student}
	return &temp
}

func printPerson(person *Person) {
	fmt.Println("Name:", person.name)
	fmt.Println("Address:", person.address)
	fmt.Println("Age:", person.age)
	if person.isStudent {
		fmt.Println("Student?: yes")
	} else {
		fmt.Println("Student?: no")
	}
}

func main() {
	person := newPerson("John", "123 Main St", 25, true)
	printPerson(person)
}
