package main

import (
	"fmt"

	"github.com/DawieMalmsteel/helloworld/cmd"
)

func concur() {
	generator := func() string {
		return "generator"
	}
	fmt.Println(generator())
	cmd.Echo("hi")
}
