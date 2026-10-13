package tasks_transport_http

import (
	"fmt"
	"net/http"

	"github.com/MaximKachkov/ToDooo/internal/core/domain"
	core_logger "github.com/MaximKachkov/ToDooo/internal/core/logger"
	core_http_request "github.com/MaximKachkov/ToDooo/internal/core/transport/http/request"
	core_http_response "github.com/MaximKachkov/ToDooo/internal/core/transport/http/response"
	core_http_types "github.com/MaximKachkov/ToDooo/internal/core/transport/http/types"
)

type PatchTaskRequest struct {
	Title       core_http_types.Nullable[string] `json:"title"`
	Description core_http_types.Nullable[string] `json:"description"`
	Completed   core_http_types.Nullable[bool]   `json:"completed"`
}
type PatchTaskResponse TaskDTOResponse

func (r *PatchTaskRequest) Validate() error {
	if r.Title.Set {
		if r.Title.Value == nil {
			return fmt.Errorf("title is set but value is nil")
		}
		if len([]rune(*r.Title.Value)) > 100 && len([]rune(*r.Title.Value)) < 1 {
			return fmt.Errorf("title should not be greater than 100")
		}
	}
	if r.Description.Set {
		if r.Description.Value != nil {
			if len([]rune(*r.Description.Value)) < 1 && len([]rune(*r.Description.Value)) > 1000 {
				return fmt.Errorf("description must be below 1000 symbols")
			}
		}
	}
	if r.Completed.Set {
		if r.Completed.Value == nil {
			return fmt.Errorf("complteted cannot be null")
		}
	}

	return nil
}

func (h *TasksHTTPHandler) PatchTask(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(logger, rw)

	taskID, err := core_http_request.GetInPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "error getting path value PathTask")
		return
	}
	var request PatchTaskRequest

	err = core_http_request.DecodeAndValidateRequest(r, &request)

	if err != nil {
		responseHandler.ErrorResponse(err, "error validating and decoding")
	}
	taskPatchDomain := DomainPatchFromRequest(request)

	taskDomain, err := h.TaskService.PatchTask(ctx, taskPatchDomain, taskID)
	if err != nil {
		responseHandler.ErrorResponse(err, "error patchTask")
	}
	taskDTO := PatchTaskResponse(DTOfromDomain(taskDomain))

	responseHandler.JSONResponse(taskDTO, http.StatusOK)

}

func DomainPatchFromRequest(req PatchTaskRequest) domain.TaskPatch {
	return domain.NewTaskPatch(req.Title.ToDomainNullable(), req.Description.ToDomainNullable(), req.Completed.ToDomainNullable())

}
