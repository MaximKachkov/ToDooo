package user_postgres_repository

import (
	"context"
	"fmt"

	core_errors "github.com/MaximKachkov/ToDooo/internal/core/errors"
)

func (s *UsersRepository) DeleteUser(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, s.pool.OpTimeout())
	defer cancel()

	query := `
	DELETE FROM todoapp.users
	WHERE id = $1;`

	commTag, err := s.pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("executing delete user query: %w", err)
	}
	if commTag.RowsAffected() == 0 {
		return fmt.Errorf("no user found with id %d: %w", id, core_errors.ErrNotFound)
	}

	return nil
}
