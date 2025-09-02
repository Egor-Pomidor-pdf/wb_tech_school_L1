package main

import "fmt"


type Human struct {
	name string
	age int
}

func (human Human) printAge() {
	fmt.Printf("%v", human.age)
}

type Action struct {
	Human
	name string
}

func main() {
	act := Action{
		Human: Human{
			name: "Egor",
			age: 19,
		},
		name: "run",
	}
	act.printAge()
}