package users_service

import (
	"context"
	"fmt"

	"github.com/MaximKachkov/ToDooo/internal/core/domain"
	core_errors "github.com/MaximKachkov/ToDooo/internal/core/errors"
)

func (s *UserService) GetUser(ctx context.Context, id int) (domain.User, error) {
	if id < 0 {
		return domain.User{}, fmt.Errorf("invalid user id:%w", core_errors.ErrInvalidArgument)
	}
	UserDomain, err := s.UsersRepository.GetUser(ctx, id)
	if err != nil {
		return domain.User{}, fmt.Errorf("No user found:%w", err)
	}
	return UserDomain, nil

}
