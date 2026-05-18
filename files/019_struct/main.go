package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func newPerson(name string) person {
	p := person{name: name}
	p.age = 18

	return p
}

func main() {
	fmt.Println(person{"Bob", 24})

	fmt.Println(person{name: "Alice", age: 36})

	fmt.Println(person{name: "Fred", age: 20})

	fmt.Println(&person{name: "Ann", age: 40})

	fmt.Println(newPerson("Osiris"))

	s := person{name: "Seaan", age: 50}
	fmt.Println(s.name)

	sptr := &s

	fmt.Println(sptr.age)

	sptr.age = 51
	fmt.Println(sptr.age)

	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}
