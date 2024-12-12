package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	name  string
	age   int
	email string
}

// Is part of a structure since we use a receptor
func (p *Person) sayHi() string {
	//fmt.Println("Hi my name is", p.name)
	return "Hi my name is " + p.name + " my age is " + strconv.Itoa(p.age) + " and my email is : " + p.email
}

func main() {
	/*
		var person Person
		person.name = "John Doe"
		person.age = 25
		person.email = "johndoe@gmail.com"
	*/
	//Another way
	person := Person{"Jhon Doe", 25, "jhon@doe.com"}
	fmt.Println(person)
	person2 := Person{"Braulio Villegas", 18, "braulio.villegas2019@gmail.com"}
	fmt.Println(person2)

	//---------------------------Pointers and methods -----------------------------
	var x int = 10
	var p *int = &x
	fmt.Println(p)
	fmt.Println(x)
	edit(&x)
	fmt.Println(x)
	person3 := &Person{"Alex", 12, "alex@gmail.com"}
	information := person3.sayHi()
	fmt.Println(information)
}

func edit(x *int) {
	*x = 20
}
