package users_transport_http

import (
	"net/http"

	core_logger "github.com/MaximKachkov/ToDooo/internal/core/logger"
	core_http_request "github.com/MaximKachkov/ToDooo/internal/core/transport/http/request"
	core_http_response "github.com/MaximKachkov/ToDooo/internal/core/transport/http/response"
)

// DeleteUser  godoc
// @Summary Удалить пользлвателя
// @Description Удалить пользователя из системы по ID
// @Tags        users
// @Param  		id path int true "ID удаляемого пользователя"
// @Success     204  "Успешное удаление пользователя"
// @Failure  	400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure  	404 {object} core_http_response.ErrorResponse "Not found"
// @Failure  	500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router      /users/{id} [delete]
func (h *UsersHTTPHandler) DeleteUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(logger, rw)

	userID, err := core_http_request.GetInPathValue(r, "id")
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
