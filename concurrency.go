package main

import "fmt"

func concur() {
	generator := func() string {
		return "generator"
	}
	fmt.Println(generator())
}
