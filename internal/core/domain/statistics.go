package domain

import "time"

type Statistics struct {
	TasksCreated              int
	TasksCompleted            int
	TaskCompletedRate         *float64
	TaskAverageComplteionTime *time.Duration
}

func NewStatistics(tasksCreated int, tasksCompleted int, taskCompletedRate *float64, taskAverageComplteionTime *time.Duration) Statistics {
	return Statistics{
		TasksCreated:              tasksCreated,
		TasksCompleted:            tasksCompleted,
		TaskCompletedRate:         taskCompletedRate,
		TaskAverageComplteionTime: taskAverageComplteionTime,
	}
}
