package user_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/MaximKachkov/ToDooo/internal/core/domain"
	core_errors "github.com/MaximKachkov/ToDooo/internal/core/errors"
	"github.com/jackc/pgx/v5"
)

func (r *UsersRepository) PatchUser(ctx context.Context, id int, patch domain.User) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	UPDATE todoapp.users
	SET
		full_name = $1,
		phone_number = $2,
		version = version + 1
	WHERE id = $3 AND version = $4	
	RETURNING id , version , full_name , phone_number;
	`
	row := r.pool.QueryRow(ctx, query, patch.FullName, patch.PhoneNumber, id, patch.Version)

	var userModel UserModel
	err := row.Scan(
		&userModel.ID,
		&userModel.Version,
		&userModel.FullName,
		&userModel.PhoneNumber,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, fmt.Errorf("no rows found :%v , %w", err, core_errors.ErrNotFound)
		}
		return domain.User{}, fmt.Errorf("error scan:%w ", err)
	}

	userReponse := domain.NewUser(userModel.ID, userModel.Version, userModel.FullName, userModel.PhoneNumber)
	return userReponse, nil
}
