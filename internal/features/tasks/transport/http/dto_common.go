package tasks_transport_http

import (
	"time"

	"github.com/MaximKachkov/ToDooo/internal/core/domain"
)

type TaskDTOResponse struct {
	ID           int        `json:"id"`
	Version      int        `json:"version"`
	Title        string     `json:"title"`
	Description  *string    `json:"desc"`
	Completed    bool       `json:"completed"`
	CreatedAt    time.Time  `json:"created_at"`
	CompletedAt  *time.Time `json:"completed_at"`
	AuthorUserID int        `json:"author_user_id"`
}

func DTOfromDomain(task domain.Task) TaskDTOResponse {
	return TaskDTOResponse{
		ID:           task.ID,
		Version:      task.Version,
		Title:        task.Title,
		Description:  task.Description,
		Completed:    task.Completed,
		CreatedAt:    task.CreatedAt,
		CompletedAt:  task.CompletedAt,
		AuthorUserID: task.AuthorUserId,
	}
}
func DTOFromDomains(tasks []domain.Task) []TaskDTOResponse {
	result := make([]TaskDTOResponse, 0, len(tasks))
	for _, task := range tasks {
		dto := DTOfromDomain(task)
		result = append(result, dto)
	}
	return result
}
