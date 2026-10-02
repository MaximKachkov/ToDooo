package user_postgres_repository

import (
	"context"
	"fmt"

	"github.com/MaximKachkov/ToDooo/internal/core/domain"
)

func (r *UsersRepository) GetUsers(ctx context.Context, limit *int, offset *int) ([]domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
SELECT id , version , full_name , phone_number
FROM todoapp.users 
ORDER BY id ASC
LIMIT $1
OFFSET $2;`

	rows, err := r.pool.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("error query retrieve: %w", err)
	}
	defer rows.Close()
	var userModels []UserModel
	for rows.Next() {
		var userModel UserModel

		err := rows.Scan(
			&userModel.ID,
			&userModel.Version,
			&userModel.FullName,
			&userModel.PhoneNumber,
		)
		if err != nil {
			return nil, fmt.Errorf("scanning usermodels error:%w", err)
		}

		userModels = append(userModels, userModel)

	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("next rows:%w", err)
	}
	userDomains := userDomainsFromModels(userModels)

	return userDomains, nil
}
