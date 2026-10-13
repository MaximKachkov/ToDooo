package tasks_postgres_repository

import (
	"time"

	"github.com/MaximKachkov/ToDooo/internal/core/domain"
)

type TaskModel struct {
	ID          int
	Version     int
	Title       string
	Description *string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time

	AuthorUserId int
}

func TaskDomainFromModel(taskModel TaskModel) domain.Task {
	return domain.NewTask(taskModel.ID, taskModel.Version, taskModel.Title, taskModel.Description, taskModel.Completed, taskModel.CreatedAt, taskModel.CompletedAt, taskModel.AuthorUserId)
}

func TaskDomainsFromModels(taskModels []TaskModel) []domain.Task {
	result := make([]domain.Task, 0, len(taskModels))
	for _, model := range taskModels {
		taskDomain := TaskDomainFromModel(model)
		result = append(result, taskDomain)
	}
	return result
}
