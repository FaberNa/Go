package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args
	fmt.Println("name of executable " + args[0])
	if len(args) > 1 {
		fmt.Println(args[1])
	} else {
		fmt.Println("Hello Faber")
	}

}
