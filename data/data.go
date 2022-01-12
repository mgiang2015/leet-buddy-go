package data

import (
	"github.com/mgiang2015/leet-buddy-go/models"
)

// Singleton design
type Database interface {
	AddProblem(p models.Problem) (bool, error)
	GetProblems() ([]models.Problem, error)
}

type databaseImpl struct {
	problems []models.Problem
}

func (db *databaseImpl) GetProblems() ([]models.Problem, error) {
	return db.problems, nil
}

func (db *databaseImpl) AddProblem(p models.Problem) (bool, error) {
	db.problems = append(db.problems, p)
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

func ClearDatabase() {
	if dbMaster == nil {
		return
	}

	dbMaster = newDatabase()
}
