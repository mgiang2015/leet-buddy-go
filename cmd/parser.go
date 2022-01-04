package cmd

import (
	"errors"
	"strings"
	"time"
)

func parseInput(input string) Command {
	var outCmd Command = nil

	switch strings.ToLower(strings.TrimSpace(input)) {
	case "create":
		// temporary
		outCmd = NewCreateCommand("Standard Url", time.Now(), false)
	case "read":
		// temporary
		outCmd = NewReadCommand()
	case "update":
		outCmd = nil
	case "delete":
		outCmd = nil
	default:
		outCmd = nil
	}

	return outCmd
}

func RunCommand(input string) (string, error) {
	currCommand := parseInput(input)

	if currCommand == nil {
		return "", errors.New("Invalid command")
	}

	output, err := currCommand.Execute()

	return output, err
}
