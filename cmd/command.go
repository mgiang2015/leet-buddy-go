package cmd

import (
	"time"

	data "github.com/mgiang2015/leet-buddy-go/data"
	"github.com/mgiang2015/leet-buddy-go/models"
)

type Command interface {
	Execute() (string, error)
}

// format:
// create -url url *required -d deadline (default 1 week from today) -done true/false (default false)
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

func newCreateCommand(url string, deadline time.Time, isDone bool) *CreateCommand {
	return &CreateCommand{
		url:      url,
		deadline: deadline,
		isDone:   isDone,
	}
}

func NewCreateCommandUrlOnly(url string) *CreateCommand {
	defDeadline := time.Now().AddDate(0, 0, 7) // default 7 days from now
	defIsDone := false

	return newCreateCommand(url, defDeadline, defIsDone)
}

func NewCreateCommandUrlDeadline(url string, dateOffset int) *CreateCommand {
	defIsDone := false
	deadline := time.Now().AddDate(0, 0, dateOffset)

	return newCreateCommand(url, deadline, defIsDone)
}

func NewReadCommand() *ReadCommand {
	return &ReadCommand{}
}
