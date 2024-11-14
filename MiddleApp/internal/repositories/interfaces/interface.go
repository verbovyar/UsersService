package interfaces

import "MiddleApp/internal/domain"

type Repository interface {
	Create(name, surname string, age uint32) error
	Update(id uint32, name, surname string, age uint32) error
	Read() []*domain.User
	Delete(id uint32) error
}
