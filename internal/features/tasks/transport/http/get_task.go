package tasks_transport_http

import (
	"net/http"

	core_logger "github.com/MaximKachkov/ToDooo/internal/core/logger"
	core_http_request "github.com/MaximKachkov/ToDooo/internal/core/transport/http/request"
	core_http_response "github.com/MaximKachkov/ToDooo/internal/core/transport/http/response"
)

type GetTaskDTOResponse TaskDTOResponse

func (h *TasksHTTPHandler) GetTask(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(logger, rw)

	taskId, err := core_http_request.GetInPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "error get task in path id")
		return
	}

	taskDomain, err := h.TaskService.GetTask(ctx, taskId)
	if err != nil {
		responseHandler.ErrorResponse(err, "error get task")
		return
	}

	taskDTO := GetTaskDTOResponse(DTOfromDomain(taskDomain))

	responseHandler.JSONResponse(taskDTO, http.StatusOK)

}
