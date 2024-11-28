package db

import (
	"MiddleApp/internal/domain"
	"MiddleApp/internal/repositories/interfaces"
	"context"
	"errors"
	"log"
)

type UsersRepository struct {
	Pool interfaces.PgxPoolIface
}

func New(pool interfaces.PgxPoolIface) *UsersRepository {
	return &UsersRepository{
		Pool: pool,
	}
}

func (r *UsersRepository) Create(name, surname string, age uint32) (error, uint32) {
	ctx := context.Background()

	user := domain.New(name, surname, age)

	query := `INSERT INTO Users (name, surname, age) VALUES ($1, $2, $3)`
	_, err := r.Pool.Query(ctx, query, user.Name, user.Surname, user.Age)
	if err != nil {
		log.Fatalf("Didnt insert user in db: %s", err.Error())

		return err, 0
	}

	var id uint32
	query = `SELECT id FROM Users WHERE name = $1`
	err = r.Pool.QueryRow(ctx, query, user.Name).Scan(&id)
	if err != nil {
		log.Fatalf("Didnt get user id in db: %s", err.Error())

		return err, 0
	}

	return nil, id
}

func (r *UsersRepository) Read() []*domain.User {
	ctx := context.Background()

	query := `SELECT id, name, surname, age FROM Users`
	rows, err := r.Pool.Query(ctx, query)
	if err != nil {
		log.Fatalf("Didnt get rows db: %s", err.Error())

		return nil
	}

	users := make([]*domain.User, 0)
	for rows.Next() {
		temp := domain.User{}
		err = rows.Scan(&temp.Id, &temp.Name, &temp.Surname, &temp.Age)
		if err != nil {
			log.Fatalf("Didnt scan rows value: %s", err.Error())

			return nil
		}

		users = append(users, &temp)
	}

	rows.Close()

	return users
}

func (r *UsersRepository) Update(id uint32, name, surname string, age uint32) (error, uint32) {
	ctx := context.Background()

	query := `SELECT id FROM Users WHERE id = $1`
	rows, err := r.Pool.Query(ctx, query, id)
	if rows == nil {
		doesNotExistErr := errors.New("user does not exist")
		return doesNotExistErr, 0
	}

	query = `UPDATE Users SET name = $1, surname = $2, age = $3 WHERE id = $4`
	_, err = r.Pool.Query(ctx, query, name, surname, age, id)
	if err != nil {
		log.Fatalf("Didnt update rows value: %s", err.Error())

		return err, 0
	}

	return nil, id
}

func (r *UsersRepository) Delete(id uint32) (error, uint32) {
	ctx := context.Background()

	query := `DELETE FROM Users WHERE Id = $1`
	_, err := r.Pool.Query(ctx, query, id)
	if err != nil {
		log.Fatalf("Didnt delete rows value: %s", err.Error())

		return err, 0
	}

	return nil, id
}
