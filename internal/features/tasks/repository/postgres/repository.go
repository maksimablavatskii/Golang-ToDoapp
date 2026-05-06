package tasks_postgres_repository

import core_postrgres_pool "github.com/maksimablavatskii/Golang-ToDoap/internal/core/repository/postgres/pool"

type TasksRepository struct {
	pool core_postrgres_pool.Pool
}

func NewTasksRepository(
	pool core_postrgres_pool.Pool,
) *TasksRepository{
	return &TasksRepository{
		pool: pool,
	}
}