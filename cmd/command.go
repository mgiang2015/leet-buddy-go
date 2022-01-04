package cmd

import (
	"time"

	data "github.com/mgiang2015/leet-buddy-go/data"
	"github.com/mgiang2015/leet-buddy-go/models"
)

type Command interface {
	Execute() (string, error)
}

type CreateCommand struct {
	url      string
	deadline time.Time
	isDone   bool
}

type ReadCommand struct{}

func (createCmd *CreateCommand) Execute() (string, error) {
	url := createCmd.url
	deadline := createCmd.deadline
	isDone := createCmd.isDone

	newProblem := models.NewProblem(url, deadline, isDone)

	_, err := data.GetDatabase().AddProblem(newProblem)

	outString := "CreateCommand executes successfully"

	if err != nil {
		outString = "Something went wrong with CreateCommand"
		return outString, err
	}

	return outString, err
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

func NewCreateCommand(url string, deadline time.Time, isDone bool) *CreateCommand {
	return &CreateCommand{
		url:      url,
		deadline: deadline,
		isDone:   isDone,
	}
}

func NewReadCommand() *ReadCommand {
	return &ReadCommand{}
}
