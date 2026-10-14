package statistics_service

import (
	"context"
	"time"

	"github.com/MaximKachkov/ToDooo/internal/core/domain"
)

type StatisticsService struct {
	statisticsRepository StatisticsRepository
}

type StatisticsRepository interface {
	GetTasks(ctx context.Context, userId *int, from *time.Time, to *time.Time) ([]domain.Task, error)
}

func NewStatisticsService(stat StatisticsRepository) *StatisticsService {
	return &StatisticsService{
		statisticsRepository: stat,
	}
}
