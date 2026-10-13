package tasks_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/MaximKachkov/ToDooo/internal/core/domain"
	core_errors "github.com/MaximKachkov/ToDooo/internal/core/errors"
	core_postgres_pool "github.com/MaximKachkov/ToDooo/internal/core/repository/postgres/pool"
)

func (r *TasksRepository) CreateTask(ctx context.Context, taskDomain domain.Task) (domain.Task, error) {

	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	INSERT INTO todoapp.tasks(title , description , completed , created_at , completed_at ,author_user_id)
	VALUES($1 , $2 , $3 , $4 , $5 , $6)
	RETURNING id , version , title , description , completed , created_at ,completed_at ,author_user_id;
	`

	row := r.pool.QueryRow(ctx, query, taskDomain.Title, taskDomain.Description, taskDomain.Completed, taskDomain.CreatedAt, taskDomain.CompletedAt, taskDomain.AuthorUserId)
	var taskModel TaskModel
	err := row.Scan(
		&taskModel.ID,
		&taskModel.Version,
		&taskModel.Title,
		&taskModel.Description,
		&taskModel.Completed,
		&taskModel.CreatedAt,
		&taskModel.CompletedAt,
		&taskModel.AuthorUserId,
	)
	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrViolatesForeignKey) {
			return domain.Task{}, fmt.Errorf("%v , user with id=%d , %w", err, taskDomain.AuthorUserId, core_errors.ErrInvalidArgument)
		}
		return domain.Task{}, fmt.Errorf("scan task error:%w", err)
	}

	taskDomain = domain.NewTask(taskModel.ID, taskModel.Version, taskModel.Title, taskModel.Description, taskModel.Completed, taskModel.CreatedAt, taskDomain.CompletedAt, taskModel.AuthorUserId)

	return taskDomain, nil

}
