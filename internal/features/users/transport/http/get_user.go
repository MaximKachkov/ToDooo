package users_transport_http

import (
	"net/http"

	core_logger "github.com/MaximKachkov/ToDooo/internal/core/logger"
	core_http_request "github.com/MaximKachkov/ToDooo/internal/core/transport/http/request"
	core_http_response "github.com/MaximKachkov/ToDooo/internal/core/transport/http/response"
)

type GetUserResponse UserDTOResponse

func (h *UsersHTTPHandler) GetUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(logger, rw)
	userID, err := core_http_request.GetInPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "error get user id from path value")
		return
	}
	user, err := h.userService.GetUser(ctx, userID)
	if err != nil {
		responseHandler.ErrorResponse(err, "error getting user")
		return
	}
	userDTO := GetUserResponse(userDTOfromDomain(user))
	responseHandler.JSONResponse(userDTO, http.StatusOK)
}
