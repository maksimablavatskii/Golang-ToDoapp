package tasks_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/maksimablavatskii/Golang-ToDoap/internal/core/domain"
	core_errors "github.com/maksimablavatskii/Golang-ToDoap/internal/core/errors"
	core_postrgres_pool "github.com/maksimablavatskii/Golang-ToDoap/internal/core/repository/postgres/pool"
)

func (r *TasksRepository) CreateTask(
	ctx context.Context,
	task domain.Task,
) (domain.Task, error){
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	INSERT INTO todoapp.tasks (title, description, completed, created_at, completed_at, author_user_id)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id, version, title, description, completed, created_at, completed_at, author_user_id;
	`

	row := r.pool.QueryRow(
		ctx, 
		query, 
		task.Title, 
		task.Description, 
		task.Completed, 
		task.CreatedAt, 
		task.CompletedAt, 
		task.AuthorUserId,
	)

	var taskModel TaskModel
	err := row.Scan(
		&taskModel.ID,
		&taskModel.Version,
		&taskModel.Title,
		&taskModel.Description,
		&taskModel.Completed,
		&taskModel.CreatedAt,
		&taskModel.CompletedAt,
		&taskModel.AuthorUserId,
	)
	if err != nil{
		if errors.Is(err, core_postrgres_pool.ErrViolatesForeignKey){
			return domain.Task{}, fmt.Errorf("%v: user with id='%d': %w", err, task.AuthorUserId, core_errors.ErrNotFound)
		}
		return domain.Task{}, fmt.Errorf("scan error: %w", err)
	}

	taskDomain := taskDomainFromModel(taskModel)

	return taskDomain, nil
}