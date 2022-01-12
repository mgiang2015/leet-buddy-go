package data

import (
	"errors"
	"time"

	"github.com/mgiang2015/leet-buddy-go/models"
)

// Singleton design
type Database interface {
	AddProblem(p models.Problem) (bool, error)
	GetProblems() ([]models.Problem, error)
	FindProblem(m models.Model) (*models.Problem, error)
	UpdateProblemDeadline(models.Model, time.Time) (bool, error)
	UpdateProblemIsDone(m models.Model, newStatus bool) (bool, error)
}

type databaseImpl struct {
	problems []models.Problem
}

func (db *databaseImpl) FindProblem(m models.Model) (*models.Problem, error) {
	problems, _ := db.GetProblems()
	index := -1

	for i, problem := range problems {
		if problem.GetModel().GetId() == m.GetId() {
			index = i
		}
	}

	if index == -1 {
		return nil, errors.New("no existing problem with given ID")
	}

	return &problems[index], nil
}

func (db *databaseImpl) GetProblems() ([]models.Problem, error) {
	return db.problems, nil
}

func (db *databaseImpl) AddProblem(p models.Problem) (bool, error) {
	db.problems = append(db.problems, p)
	return true, nil
}

func (db *databaseImpl) UpdateProblemDeadline(m models.Model, newDeadline time.Time) (bool, error) {
	problem, err := db.FindProblem(m)
	if err != nil {
		return false, err
	}

	if time.Now().After(newDeadline) {
		return false, errors.New("given deadline has already passed")
	}

	if problem.GetDeadline().Equal(newDeadline) {
		return false, errors.New("given deadline is the same as current")
	}

	problem.SetDeadline(newDeadline)
	return true, nil
}

func (db *databaseImpl) UpdateProblemIsDone(m models.Model, newStatus bool) (bool, error) {
	problem, err := db.FindProblem(m)
	if err != nil {
		return false, err
	}

	if problem.GetIsDone() == newStatus {
		return false, errors.New("problem is already done")
	}

	problem.SetIsDone(newStatus)
	return true, nil
}

func newDatabase() *databaseImpl {
	db := &databaseImpl{
		problems: make([]models.Problem, 0),
	}

	return db
}

var dbMaster Database = nil

// This will be called by clients to access the only instance of database
func GetDatabase() Database {
	if dbMaster == nil {
		dbMaster = newDatabase()
	}

	return dbMaster
}
