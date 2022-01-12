package models

type Model interface {
	GetId() int
	SetId(id int)
}

type modelImpl struct {
	id int
}

func (m *modelImpl) SetId(id int) {
	m.id = id
}

func (m *modelImpl) GetId() int {
	return m.id
}

// Temporary ID mechanism.
// TODO: Hash ID generation
var currentId int = 0

func NewModel() Model {
	modelObj := &modelImpl{
		id: currentId,
	}

	currentId++

	return modelObj
}

func NewModelWithId(id int) Model {
	modelObj := &modelImpl{
		id: id,
	}

	return modelObj
}
