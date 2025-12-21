package main

import (
	"fmt"

	"github.com/DawieMalmsteel/helloworld/fighter"
)

func main() {
	// f := fighter.Goku
	f := &fighter.Goku
	fmt.Println(f.GetName())
	fmt.Println(f.GetAge())
	f.ChangeName("Vegeta")
	f.ChangeAge(30)
	fmt.Println(f.GetName())
	fmt.Println(f.GetAge())

	fmt.Println("After change")
	fmt.Println(fighter.Goku.GetName())
	fmt.Println(fighter.Goku.GetAge())
}
