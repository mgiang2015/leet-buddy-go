package cmd

import (
	"errors"
	"strings"
)

func parseInput(input string) Command {
	var outCmd Command = nil

	params := strings.Fields(strings.ToLower(input)) // split by whitespace into fields

	switch params[0] {
	case "create":
		outCmd = CreateCommandFromParams(params)
	case "read":
		// temporary
		outCmd = NewReadCommand()
	case "update":
		outCmd = UpdateCommandFromParams(params)
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
		return "", errors.New("invalid command")
	}

	output, err := currCommand.Execute()

	return output, err
}
