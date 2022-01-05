package cmd

import data "github.com/mgiang2015/leet-buddy-go/data"

type ReadCommand struct{}

func NewReadCommand() *ReadCommand {
	return &ReadCommand{}
}

func (readCmd *ReadCommand) Execute() (string, error) {
	problems, err := data.GetDatabase().GetProblems()
	result := ""

	if err != nil {
		result = "Something wrong with ReadCommand"
		return result, err
	}

	for _, p := range problems {
		result += p.String()
	}

	return result, nil
}
