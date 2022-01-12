package cmd

import (
	"strconv"
	"time"

	"github.com/mgiang2015/leet-buddy-go/data"
	"github.com/mgiang2015/leet-buddy-go/models"
)

type UpdateCommand struct {
	model    models.Model
	deadline time.Time
	isDone   bool
}

func newUpdateCommand(model models.Model, deadline time.Time, isDone bool) *UpdateCommand {
	return &UpdateCommand{
		model:    model,
		deadline: deadline,
		isDone:   isDone,
	}
}

func UpdateCommandFromParams(params []string) *UpdateCommand {
	// -problem: Problem id model
	// -deadline: deadline
	// -isDone: isDone done

	modelIndex := 0
	deadlineIndex := 0
	isDoneIndex := 0

	var outCmd *UpdateCommand = nil

	// find parameters
	for index, param := range params {
		if param == "-isDone" {
			isDoneIndex = index + 1
		} else if param == "-deadline" {
			deadlineIndex = index + 1
		} else if param == "-problem" {
			modelIndex = index + 1
		}
	}

	if modelIndex == 0 || isDoneIndex == 0 && deadlineIndex == 0 {
		return nil
	}

	modelId, err := strconv.Atoi(params[modelIndex])
	if err != nil {
		return nil
	}

	// First find the problem the user is updating
	problem, err := data.GetDatabase().FindProblem(models.NewModelWithId(modelId))
	if err != nil {
		return nil
	}
	// set the new parameters
	var newDeadline time.Time
	var newIsDone bool
	//

	return outCmd
}

func (update *UpdateCommand) Execute() (string, error) {
	_, deadlineErr := data.GetDatabase().UpdateProblemDeadline(update.model, update.deadline)
	if deadlineErr != nil {
		return "Update problem failed", deadlineErr
	}

	_, isDoneErr := data.GetDatabase().UpdateProblemIsDone(update.model, update.isDone)
	if isDoneErr != nil {
		return "Update problem failed", isDoneErr
	}

	return "Updated problem successfully", nil
}
