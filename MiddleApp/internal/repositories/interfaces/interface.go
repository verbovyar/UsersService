package interfaces

import "MiddleApp/internal/domain"

//go:generate mockgen -source=interface.go -destination=mock/mock.go

type Repository interface {
	Create(name, surname string, age uint32) (error, uint32)
	Update(id uint32, name, surname string, age uint32) (error, uint32)
	Read() []*domain.User
	Delete(id uint32) (error, uint32)
}
