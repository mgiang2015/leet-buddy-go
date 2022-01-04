package main

import (
	"bufio"
	"fmt"
	"os"

	cmd "github.com/mgiang2015/leet-buddy-go/cmd"
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

		output, err := cmd.RunCommand(input)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		fmt.Println(output)
	}
}
