package main

import (
	"fmt"
	"io/ioutil"
	"os"

	pineapple "pineapple/src"
)

func main() {

	// read file
	args := os.Args
	if len(args) != 2 {
		fmt.Printf("Usage: %s filename\n", args[0])
		return
	}
	filename := args[1]
	code, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %s\n", filename)
		return
	}

	// execute
	pineapple.Execute(string(code))

}
