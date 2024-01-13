package model

type SimpleModel interface {
	Id() int
	setId(id int)
}

type simpleModelImpl struct {
	id int64
}

func (m *simpleModelImpl) Id() int64 {
	return m.id
}

func (m *simpleModelImpl) SetId(id int64) {
	m.id = id
}
