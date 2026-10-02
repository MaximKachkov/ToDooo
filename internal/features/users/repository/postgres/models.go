package user_postgres_repository

import "github.com/MaximKachkov/ToDooo/internal/core/domain"

type UserModel struct {
	ID      int
	Version int

	FullName    string
	PhoneNumber *string
}

func userDomainsFromModels(users []UserModel) []domain.User {
	userDomains := make([]domain.User, len(users))
	for i, user := range users {
		userDomains[i] = domain.NewUser(user.ID, user.Version, user.FullName, user.PhoneNumber)
	}
	return userDomains
}

func userDomainFromModel(user UserModel) domain.User {
	return domain.NewUser(user.ID , user.Version , user.FullName , user.PhoneNumber)
}
