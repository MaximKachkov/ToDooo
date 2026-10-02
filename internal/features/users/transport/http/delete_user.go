package users_transport_http

import (
	"net/http"

	core_logger "github.com/MaximKachkov/ToDooo/internal/core/logger"
	core_http_response "github.com/MaximKachkov/ToDooo/internal/core/transport/http/response"
	core_http_utils "github.com/MaximKachkov/ToDooo/internal/core/transport/http/utils"
)

func (h *UsersHTTPHandler) DeleteUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(logger, rw)

	userID, err := core_http_utils.GetInPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "error get user id from path value")
		return
	}

	err = h.userService.DeleteUser(ctx, userID)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to deleted user")
		return
	}
	responseHandler.NoContentResponse()

}
