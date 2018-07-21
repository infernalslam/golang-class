package main

import (
	"fmt"
)

type Person struct {
	name       string
	age        int
	programmer bool
}

type role interface {
	check()
}

func (person *Person) check() bool {
	return person.programmer == true
}

func main() {
	person := Person{"tak", 24, true}
	fmt.Println(person.name)
	fmt.Println(person.age)
	fmt.Println(person.programmer)
	fmt.Print("person role interface ::: ")
	fmt.Println(person.check())

}
