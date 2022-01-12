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

func (p *Problem) GetModel() Model {
	return p.model
}

func (p *Problem) SetDeadline(newDeadline time.Time) {
	p.deadline = newDeadline
}

func (p *Problem) GetIsDone() bool {
	return p.isDone
}

func (p *Problem) GetDeadline() time.Time {
	return p.deadline
}

func (p *Problem) SetIsDone(isDone bool) {
	p.isDone = isDone
}

func (p Problem) String() string {
	return fmt.Sprintf("%d: Url: %s, deadline: %v, isDone: %v\n", p.model.GetId(), p.url, p.deadline, p.isDone)
}
