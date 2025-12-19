package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	// code 1: chạy được
	var person *Person
	person = &Person{Name: "Alice", Age: 30}

	// code 2: đéo chạy được nil pointer dereference
	var person *Person
	*person = Person{Name: "Alice", Age: 30}

	// code 3: chạy được fix của code 2 cái hàm new nó cấp bộ nhớ cho variable luôn
	person := new(Person)
	*person = Person{Name: "Alice", Age: 30}

	fmt.Printf("%+v\n", *person)
}
