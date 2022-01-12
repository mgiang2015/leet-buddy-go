package cmd

import (
	"strconv"
	"time"

	data "github.com/mgiang2015/leet-buddy-go/data"
	"github.com/mgiang2015/leet-buddy-go/models"
)

// format:
// create -url url *required -d deadline (default 1 week from today) -done true/false (default false)
type CreateCommand struct {
	url      string
	deadline time.Time
	isDone   bool
}

func newCreateCommand(url string, deadline time.Time, isDone bool) *CreateCommand {
	return &CreateCommand{
		url:      url,
		deadline: deadline,
		isDone:   isDone,
	}
}

func newCreateCommandUrlOnly(url string) *CreateCommand {
	defDeadline := time.Now().AddDate(0, 0, 7) // default 7 days from now
	defIsDone := false

	return newCreateCommand(url, defDeadline, defIsDone)
}

func newCreateCommandUrlDeadline(url string, dateOffset int) *CreateCommand {
	defIsDone := false
	deadline := time.Now().AddDate(0, 0, dateOffset)

	return newCreateCommand(url, deadline, defIsDone)
}

func CreateCommandFromParams(params []string) *CreateCommand {
	// -url: url
	// -deadline: deadline
	urlIndex := 0
	deadlineIndex := 0

	var outCmd *CreateCommand = nil

	// find parameters
	for index, param := range params {
		if param == "-url" {
			urlIndex = index + 1
		} else if param == "-deadline" {
			deadlineIndex = index + 1
		}
	}

	if urlIndex == 0 {
		outCmd = nil
	} else if deadlineIndex == 0 {
		outCmd = newCreateCommandUrlOnly(params[urlIndex])
	} else {
		dateOffset, err := strconv.Atoi(params[deadlineIndex])
		if err != nil {
			outCmd = nil
		} else {
			outCmd = newCreateCommandUrlDeadline(params[urlIndex], dateOffset)
		}
	}
	return outCmd
}

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
