package tasks_transport_http

import (
	"time"

	"github.com/maksimablavatskii/Golang-ToDoap/internal/core/domain"
)

type TaskDTOResponse struct {
	ID      int `json:"id"`
	Version int `json:"version"`

	Title       string     `json:"title"`
	Description *string    `json:"description"`
	Completed   bool       `json:"completed"`
	CreatedAt   time.Time  `json:"created_at"`
	CompletedAt *time.Time `json:"completed_at"`

	AuthorUserId int `json:"authorUserId"`
}

func taskDTOFromDomain(task domain.Task) TaskDTOResponse {
	return TaskDTOResponse{
		ID:      task.ID,
		Version: task.Version,

		Title:       task.Title,
		Description: task.Description,
		Completed:   task.Completed,
		CreatedAt:   task.CreatedAt,
		CompletedAt: task.CompletedAt,

		AuthorUserId: task.AuthorUserId,
	}
}

func tasksDTOFromdomains(tasks []domain.Task) []TaskDTOResponse {
	tasksDTO := make([]TaskDTOResponse, len(tasks))

	for i, task := range tasks {
		tasksDTO[i] = taskDTOFromDomain(task)
	}

	return tasksDTO
}
