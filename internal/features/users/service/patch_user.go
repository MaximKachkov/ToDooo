package users_service

import (
	"context"
	"fmt"

	"github.com/MaximKachkov/ToDooo/internal/core/domain"
)

func (s *UserService) PatchUser(ctx context.Context, id int, patch domain.UserPatch) (domain.User, error) {
	user, err := s.UsersRepository.GetUser(ctx, id)
	if err != nil {
		return domain.User{}, fmt.Errorf("error getting user:%w", err)
	}

	if err := user.ApplyPatch(patch); err != nil {
		return domain.User{}, fmt.Errorf("error applying patch:%w", err)
	}

	updatedUser, err := s.UsersRepository.PatchUser(ctx, id, user)
	if err != nil {
		return domain.User{}, fmt.Errorf("error patch user in repo:%w", err)
	}
	return updatedUser, nil

}
