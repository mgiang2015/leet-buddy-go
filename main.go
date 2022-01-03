package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for { // go infinite while loop
		fmt.Print(">> ")
		// read input
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		fmt.Print("Received input: " + input)
	}
}
