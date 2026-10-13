package tasks_service

import (
	"context"
	"fmt"

	"github.com/MaximKachkov/ToDooo/internal/core/domain"
	core_errors "github.com/MaximKachkov/ToDooo/internal/core/errors"
)

func (s *TasksService) GetTasks(ctx context.Context, offset *int, limit *int, userId *int) ([]domain.Task, error) {
	if limit != nil && *limit < 0 {
		return nil, fmt.Errorf("limit must be non negative: %w", core_errors.ErrInvalidArgument)
	}

	if offset != nil && *offset < 0 {
		return nil, fmt.Errorf("offset must be non negative: %w", core_errors.ErrInvalidArgument)
	}
	if userId != nil && *userId < 0 {
		return nil, fmt.Errorf("offset must be non negative: %w", core_errors.ErrInvalidArgument)
	}

	tasksDomain, err := s.TasksRepository.GetTasks(ctx, offset, limit, userId)
	if err != nil {
		return nil, fmt.Errorf("get task Repository error :%w", err)
	}
	return tasksDomain, nil

}
