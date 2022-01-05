package cmd

import (
	"errors"
	"strconv"
	"strings"
)

func parseInput(input string) Command {
	var outCmd Command = nil

	fields := strings.Fields(strings.ToLower(input)) // split by whitespace into fields

	switch fields[0] {
	case "create":
		// -url: url
		// -deadline: deadline
		urlIndex := 0
		deadlineIndex := 0

		// find parameters
		for index, param := range fields {
			if param == "-url" {
				urlIndex = index + 1
			} else if param == "-deadline" {
				deadlineIndex = index + 1
			}
		}

		if urlIndex == 0 {
			outCmd = nil
		} else if deadlineIndex == 0 {
			outCmd = NewCreateCommandUrlOnly(fields[urlIndex])
		} else {
			dateOffset, err := strconv.Atoi(fields[deadlineIndex])
			if err != nil {
				outCmd = nil
			} else {
				outCmd = NewCreateCommandUrlDeadline(fields[urlIndex], dateOffset)
			}
		}

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
		return "", errors.New("invalid command")
	}

	output, err := currCommand.Execute()

	return output, err
}
