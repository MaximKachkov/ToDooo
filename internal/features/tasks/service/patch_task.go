package tasks_service

import (
	"context"
	"fmt"

	"github.com/MaximKachkov/ToDooo/internal/core/domain"
)

func (s *TasksService) PatchTask(ctx context.Context, patch domain.TaskPatch, id int) (domain.Task, error) {
	task, err := s.TasksRepository.GetTask(ctx, id)
	if err != nil {
		return domain.Task{}, fmt.Errorf("error while getting task:%w", err)
	}

	err = task.ApplyPatch(patch)
	if err != nil {
		return domain.Task{}, fmt.Errorf("error while apply patch:%w", err)
	}

	taskDomain, err := s.TasksRepository.PatchTask(ctx, task, id)
	if err != nil {
		return domain.Task{}, fmt.Errorf("patchTask error:%w", err)
	}

	return taskDomain, nil

}
