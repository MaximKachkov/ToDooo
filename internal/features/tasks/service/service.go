package tasks_service

import (
	"context"

	"github.com/MaximKachkov/ToDooo/internal/core/domain"
)

type TasksService struct {
	TasksRepository
}

type TasksRepository interface {
	CreateTask(ctx context.Context, taskDomain domain.Task) (domain.Task, error)
	GetTasks(ctx context.Context, offset *int, limit *int, userId *int) ([]domain.Task, error)
	GetTask(ctx context.Context, id int) (domain.Task, error)
	DeleteTask(ctx context.Context, id int) error
	PatchTask(ctx context.Context, taskDomain domain.Task, id int) (domain.Task, error)
}

func NewTasksService(tr TasksRepository) *TasksService {
	return &TasksService{
		TasksRepository: tr,
	}
}
