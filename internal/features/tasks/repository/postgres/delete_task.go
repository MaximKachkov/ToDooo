package tasks_postgres_repository

import (
	"context"
	"fmt"

	core_errors "github.com/MaximKachkov/ToDooo/internal/core/errors"
)

func (r *TasksRepository) DeleteTask(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	DELETE from todoapp.tasks
	WHERE id = $1
	`
	commTag, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("error executing a query : %w", err)
	}
	if commTag.RowsAffected() == 0 {
		return fmt.Errorf("no row with this id found %d: %w ", id, core_errors.ErrNotFound)
	}
	return nil
}
