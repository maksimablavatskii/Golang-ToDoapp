package users_postgres_repository

import (
	core_postrgres_pool "github.com/maksimablavatskii/Golang-ToDoap/internal/core/repository/postgres/pool"
)

type UsersRepository struct {
	pool core_postrgres_pool.Pool
}

func NewUsersRepository(pool core_postrgres_pool.Pool) *UsersRepository {
	return &UsersRepository{
		pool: pool,
	}
}
