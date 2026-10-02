package users_service

import (
	"context"
	"fmt"

	core_errors "github.com/MaximKachkov/ToDooo/internal/core/errors"
)

func (s *UserService) DeleteUser(ctx context.Context, id int) error {
	if id < 0 {
		return fmt.Errorf("invalid user id:%w", core_errors.ErrInvalidArgument)
	}
	err := s.UsersRepository.DeleteUser(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete user:%w", err)
	}
	return nil
}
