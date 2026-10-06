package users_transport_http

import (
	"context"
	"net/http"

	"github.com/MaximKachkov/ToDooo/internal/core/domain"
	core_http_server "github.com/MaximKachkov/ToDooo/internal/core/transport/http/server"
)

type UsersHTTPHandler struct {
	userService UserService
}

type UserService interface {
	CreateUser(
		ctx context.Context,
		user domain.User,
	) (domain.User, error)
	GetUsers(
		ctx context.Context,
		limit *int,
		offset *int,
	) ([]domain.User, error)
	GetUser(ctx context.Context, id int) (domain.User, error)
	DeleteUser(ctx context.Context, id int) error
	PatchUser(ctx context.Context, id int, patch domain.UserPatch) (domain.User, error)
}

func NewUsersHTTPHandler(us UserService) *UsersHTTPHandler {
	return &UsersHTTPHandler{
		userService: us,
	}
}

func (u *UsersHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/users",
			Handler: u.CreateUser,
		},
		{
			Method:  http.MethodGet,
			Path:    "/users",
			Handler: u.GetUsers,
			//Middleware: []core_http_middleware.Middleware{
			//	core_http_middleware.Dummy("xui"),
			//},
		},
		{
			Method:  http.MethodGet,
			Path:    "/users/{id}",
			Handler: u.GetUser,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/users/{id}",
			Handler: u.DeleteUser,
		},
		{
			Method:  http.MethodPatch,
			Path:    "/users/{id}",
			Handler: u.PatchUser,
		},
	}
}
