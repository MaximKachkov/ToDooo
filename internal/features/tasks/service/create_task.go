package tasks_service

import (
	"context"
	"fmt"

	"github.com/MaximKachkov/ToDooo/internal/core/domain"
)

func (s *TasksService) CreateTask(ctx context.Context, taskDomain domain.Task) (domain.Task, error) {
	if err := taskDomain.Validate(); err != nil {
		return domain.Task{}, err
	}

	taskDomain, err := s.TasksRepository.CreateTask(ctx, taskDomain)
	if err != nil {
		return domain.Task{}, fmt.Errorf("error while creating a task:%w", err)
	}
	return taskDomain, nil

}
