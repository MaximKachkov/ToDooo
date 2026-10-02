package users_service

import (
	"context"
	"fmt"

	"github.com/MaximKachkov/ToDooo/internal/core/domain"
	core_errors "github.com/MaximKachkov/ToDooo/internal/core/errors"
)

func (s *UserService) GetUsers(ctx context.Context, limit *int, offset *int) ([]domain.User, error) {
	if limit != nil && *limit < 0 {
		return nil, fmt.Errorf("limit must be non negative: %w", core_errors.ErrInvalidArgument)
	}

	if offset != nil && *offset < 0 {
		return nil, fmt.Errorf("offset must be non negative: %w", core_errors.ErrInvalidArgument)
	}

	userDomains, err := s.UsersRepository.GetUsers(ctx, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("problem getting user domains:%w", err)
	}
	return userDomains, nil
}
