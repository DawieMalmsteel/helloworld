package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	// concur()

	mapPersons := map[string]string{
		"Alice":      "30",
		"Bob":        "25",
		"Borderland": "40",
	}

	for i, p := range mapPersons {
		fmt.Println(i, p)
	}
}
