package tasks_transport_http

import (
	"context"
	"net/http"

	"github.com/MaximKachkov/ToDooo/internal/core/domain"
	core_http_server "github.com/MaximKachkov/ToDooo/internal/core/transport/http/server"
)

type TasksHTTPHandler struct {
	TaskService
}

type TaskService interface {
	CreateTask(ctx context.Context, taskDomain domain.Task) (domain.Task, error)
	GetTasks(ctx context.Context, offset *int, limit *int, userId *int) ([]domain.Task, error)
	GetTask(ctx context.Context, id int) (domain.Task, error)
	DeleteTask(ctx context.Context, id int) error
	PatchTask(ctx context.Context, patch domain.TaskPatch, id int) (domain.Task, error)
}

func NewTasksHTTPHandler(ts TaskService) *TasksHTTPHandler {
	return &TasksHTTPHandler{
		TaskService: ts,
	}
}

func (u *TasksHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/tasks",
			Handler: u.CreateTask,
		},
		{
			Method:  http.MethodGet,
			Path:    "/tasks",
			Handler: u.GetTasks,
		},
		{
			Method:  http.MethodGet,
			Path:    "/tasks/{id}",
			Handler: u.GetTask,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/tasks/{id}",
			Handler: u.DeleteTask,
		},
		{
			Method:  http.MethodPatch,
			Path:    "/tasks/{id}",
			Handler: u.PatchTask,
		},
	}
}
