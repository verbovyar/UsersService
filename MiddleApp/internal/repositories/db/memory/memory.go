package memory

import (
	"MiddleApp/internal/domain"
	"MiddleApp/internal/repositories/interfaces"
	"log"
)

var LastId = uint32(0)

type Memory struct {
	Data map[uint32]*domain.User

	repo interfaces.Repository
}

func New(data map[uint32]*domain.User) *Memory {
	return &Memory{Data: data}
}

func (m *Memory) Create(name, surname string, age uint32) (error, uint32) {
	user := &domain.User{
		Name:    name,
		Surname: surname,
		Age:     age,
		Id:      LastId,
	}

	m.Data[LastId] = user

	LastId++

	return nil, user.Id
}

func (m *Memory) Read() []*domain.User {
	list := make([]*domain.User, len(m.Data))

	for i, val := range m.Data {
		list[i] = &domain.User{
			Name:    val.Name,
			Surname: val.Surname,
			Age:     val.Age,

			Id: val.Id,
		}
	}

	return list
}

func (m *Memory) Update(id uint32, name, surname string, age uint32) (error, uint32) {
	m.Data[id].Name = name
	m.Data[id].Surname = surname
	m.Data[id].Age = age

	return nil, id
}

func (m *Memory) Delete(id uint32) (error, uint32) {
	if _, ok := m.Data[id]; ok {
		delete(m.Data, id)
	} else {
		log.Println("Delete user error")
	}

	return nil, id
}
