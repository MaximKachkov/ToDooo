package statistics_service

import (
	"context"
	"fmt"
	"time"

	"github.com/MaximKachkov/ToDooo/internal/core/domain"
	core_errors "github.com/MaximKachkov/ToDooo/internal/core/errors"
)

func (s *StatisticsService) GetStatistics(ctx context.Context, id *int, from *time.Time, to *time.Time) (domain.Statistics, error) {
	if from != nil && to != nil {
		if to.Before(*from) || to.Equal(*from) {
			return domain.Statistics{}, fmt.Errorf("from cannot be before to :%s , %s , %w", from.String(), to.String(), core_errors.ErrInvalidArgument)
		}
	}
	tasks, err := s.statisticsRepository.GetTasks(ctx, id, from, to)
	if err != nil {
		return domain.Statistics{}, fmt.Errorf("error get tasks:%w", err)
	}
	statistics := CalcStatistics(tasks)
	return statistics, nil

}

func CalcStatistics(tasks []domain.Task) domain.Statistics {
	if len(tasks) == 0 {
		return domain.Statistics{
			TasksCreated:              0,
			TasksCompleted:            0,
			TaskCompletedRate:         nil,
			TaskAverageComplteionTime: nil,
		}
	}
	tasksCreated := len(tasks)
	tasksCompleted := 0
	var totalCompletedDuration time.Duration
	for _, task := range tasks {
		if task.Completed == true {
			tasksCompleted++
		}
		taskCompletionDuration := task.CompletionDuration()
		if taskCompletionDuration != nil {
			totalCompletedDuration += *taskCompletionDuration
		}
	}
	var TasksAverageCompletionTime *time.Duration
	if tasksCompleted > 0 && totalCompletedDuration != 0 {
		avg := totalCompletedDuration / time.Duration(tasksCompleted)
		TasksAverageCompletionTime = &avg
	}

	taskCompltetedRate := float64(tasksCompleted) / float64(tasksCreated) * 100

	return domain.NewStatistics(tasksCreated, tasksCompleted, &taskCompltetedRate, TasksAverageCompletionTime)

}
