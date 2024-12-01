package domain

type User struct {
	Id uint32

	Name    string
	Surname string
	Age     uint32
}

func New(name, surname string, age uint32) *User {
	user := User{
		Name:    name,
		Surname: surname,
		Age:     age,
	}

	return &user
}
