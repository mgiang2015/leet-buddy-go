package models

import (
	"fmt"
	"time"
)

type Problem struct {
	model    Model
	url      string
	deadline time.Time
	isDone   bool
}

func NewProblem(url string, dealine time.Time, isDone bool) Problem {
	newModel := NewModel()
	newProblem := Problem{
		model:    newModel,
		url:      url,
		deadline: dealine,
		isDone:   isDone,
	}

	return newProblem
}

func (p Problem) String() string {
	return fmt.Sprintf("Url: %s, deadline: %v, isDone: %v\n", p.url, p.deadline, p.isDone)
}
