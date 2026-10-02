package users_transport_http

import "github.com/MaximKachkov/ToDooo/internal/core/domain"

type UserDTOResponse struct {
	ID          int     `json:"id"`
	Version     int     `json:"version"`
	FullName    string  `json:"full_name"`
	PhoneNumber *string `json:"phone_number"`
}

func userDTOfromDomain(user domain.User) UserDTOResponse {
	return UserDTOResponse{
		ID:          user.ID,
		Version:     user.Version,
		FullName:    user.FullName,
		PhoneNumber: user.PhoneNumber,
	}
}

func usersDTOfromDomains(users []domain.User) []UserDTOResponse {
	usersDTO := make([]UserDTOResponse, len(users))

	for i, user := range users {
		domainUser := userDTOfromDomain(user)
		usersDTO[i] = domainUser
	}
	return usersDTO
}
