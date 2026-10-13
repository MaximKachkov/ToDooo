package tasks_service

import (
	"context"
	"fmt"

	core_errors "github.com/MaximKachkov/ToDooo/internal/core/errors"
)

func (s *TasksService) DeleteTask(ctx context.Context, id int) error {

	if id < 0 {
		return fmt.Errorf("id should be greater than 0 :%w", core_errors.ErrInvalidArgument)
	}

	err := s.TasksRepository.DeleteTask(ctx, id)
	if err != nil {
		return fmt.Errorf("delete task error : %w", err)
	}
	return nil
}
