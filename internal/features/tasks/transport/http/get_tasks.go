package tasks_transport_http

import (
	"fmt"
	"net/http"

	core_logger "github.com/MaximKachkov/ToDooo/internal/core/logger"
	core_http_request "github.com/MaximKachkov/ToDooo/internal/core/transport/http/request"
	core_http_response "github.com/MaximKachkov/ToDooo/internal/core/transport/http/response"
)

type TasksDtoResponses []TaskDTOResponse

func (h *TasksHTTPHandler) GetTasks(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(logger, rw)

	offset, limit, userId, err := GetUserIDLimitOffsetQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "error getTasks get limit off query params")
		return
	}

	tasksDomain, err := h.TaskService.GetTasks(ctx, offset, limit, userId)
	if err != nil {
		responseHandler.ErrorResponse(err, " getTasks error")
		return
	}
	tasksDTOs := TasksDtoResponses(DTOFromDomains(tasksDomain))
	responseHandler.JSONResponse(tasksDTOs, http.StatusOK)

}

func GetUserIDLimitOffsetQueryParams(r *http.Request) (*int, *int, *int, error) {
	const (
		limitConst  string = "limit"
		offsetConst string = "offset"
		userIdConst string = "user_id"
	)
	offset, err := core_http_request.GetIntQueryParam(r, offsetConst)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get offset params:%w", err)
	}
	limit, err := core_http_request.GetIntQueryParam(r, limitConst)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get limit params:%w", err)
	}
	userId, err := core_http_request.GetIntQueryParam(r, userIdConst)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get userid params:%w", err)
	}

	return offset, limit, userId, nil
}
