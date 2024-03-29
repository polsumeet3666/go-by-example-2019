package main

import "fmt"

type person struct {
	name string
	age  int
}

func NewPerson(name string) *person {
	p := person{name: name}
	p.age = 42

	return &p
}

func main() {
	fmt.Println(person{"bob", 20})

	fmt.Println(person{name: "alice", age: 30})

	fmt.Println(person{name: "fred"})

	fmt.Println(&person{name: "Ann", age: 50})

	fmt.Println(NewPerson("jon"))

	s := person{name: "sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(s.age)

}
