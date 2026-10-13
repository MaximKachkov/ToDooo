package tasks_service

import (
	"context"
	"fmt"

	"github.com/MaximKachkov/ToDooo/internal/core/domain"
	core_errors "github.com/MaximKachkov/ToDooo/internal/core/errors"
)

func (s *TasksService) GetTask(ctx context.Context, id int) (domain.Task, error) {
	if id < 0 {
		return domain.Task{}, fmt.Errorf("id should be greater than 0:%w", core_errors.ErrInvalidArgument)
	}
	taskDomain, err := s.TasksRepository.GetTask(ctx, id)
	if err != nil {
		return domain.Task{}, err
	}
	return taskDomain, nil

}
