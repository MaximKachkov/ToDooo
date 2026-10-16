package users_transport_http

import "github.com/MaximKachkov/ToDooo/internal/core/domain"

type UserDTOResponse struct {
	ID          int     `json:"id"                  example:"10"`
	Version     int     `json:"version"             example:"3"`
	FullName    string  `json:"full_name"           example:"Максим Скачков"`
	PhoneNumber *string `json:"phone_number"        example:"+72281540727"`
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
