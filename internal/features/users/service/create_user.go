package users_service

import (
	"context"
	"fmt"

	"github.com/MaximKachkov/ToDooo/internal/core/domain"
)

func (s *UserService) CreateUser(
	ctx context.Context,
	user domain.User,
) (domain.User, error) {
	if err := user.Validate(); err != nil {
		return domain.User{}, fmt.Errorf("create user validation erorr : %w", err)
	}
	user, err := s.UsersRepository.CreateUser(ctx, user)
	if err != nil {
		return domain.User{}, fmt.Errorf("create user repository: %w", err)
	}
	return user, nil
}
