package users_transport_http

import (
	"fmt"
	"net/http"

	"github.com/MaximKachkov/ToDooo/internal/core/domain"
	core_logger "github.com/MaximKachkov/ToDooo/internal/core/logger"
	core_http_request "github.com/MaximKachkov/ToDooo/internal/core/transport/http/request"
	core_http_response "github.com/MaximKachkov/ToDooo/internal/core/transport/http/response"
	"go.uber.org/zap"
)

type GetUsersResponse []UserDTOResponse

// GetUsers   godoc
// @Summary  Получить пользователей
// @Description Получить пользователей с лимит и оффсет
// @Tags users
// @Produce json
// @Param limit query int false "size of the page with users"
// @Param offset query int false "size of the page with users"
// @Success 200 {object}  GetUsersResponse "List of users"
// @Failure 400 {object}  core_http_response.ErrorResponse "Bad Request"
// @Failure 500 {object}  core_http_response.ErrorResponse "Internal server error"
// @Router  /users [get]
func (h *UsersHTTPHandler) GetUsers(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(logger, rw)
	limit, offset, err := GetLimitOffsetQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get query params")
		return
	}
	var domainUsers []domain.User
	domainUsers, err = h.userService.GetUsers(ctx, limit, offset)
	if err != nil {
		logger.Error("wasn't able to get users list", zap.Error(err))
		responseHandler.ErrorResponse(err, "wasn't able to get users list")
	}
	response := GetUsersResponse(usersDTOfromDomains(domainUsers))

	responseHandler.JSONResponse(response, http.StatusOK)
}

func GetLimitOffsetQueryParams(r *http.Request) (*int, *int, error) {
	const (
		limitConst  string = "limit"
		offsetConst string = "offset"
	)
	offset, err := core_http_request.GetIntQueryParam(r, offsetConst)
	if err != nil {
		return nil, nil, fmt.Errorf("get offset params:%w", err)
	}
	limit, err := core_http_request.GetIntQueryParam(r, limitConst)
	if err != nil {
		return nil, nil, fmt.Errorf("get limit params:%w", err)
	}
	return limit, offset, nil
}
