package user_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/MaximKachkov/ToDooo/internal/core/domain"
	core_errors "github.com/MaximKachkov/ToDooo/internal/core/errors"
	"github.com/jackc/pgx/v5"
)

func (s *UsersRepository) GetUser(ctx context.Context, id int) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, s.pool.OpTimeout())
	defer cancel()

	query := `
   SELECT id , version , full_name , phone_number
   FROM todoapp.users
   WHERE id = $1;`

	row := s.pool.QueryRow(ctx, query, id)

	var userModel UserModel
	err := row.Scan(
		&userModel.ID,
		&userModel.Version,
		&userModel.FullName,
		&userModel.PhoneNumber,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, fmt.Errorf("problem scanning the user %d for db:%w", id, core_errors.ErrNotFound)
		}
		return domain.User{}, fmt.Errorf("problem scanning :%w", err)
	}

	userDomain := userDomainFromModel(userModel)

	return userDomain, nil

}
