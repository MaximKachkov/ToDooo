package users_transport_http

import (
	"net/http"

	"github.com/MaximKachkov/ToDooo/internal/core/domain"
	core_logger "github.com/MaximKachkov/ToDooo/internal/core/logger"
	core_http_request "github.com/MaximKachkov/ToDooo/internal/core/transport/http/request"
	core_http_response "github.com/MaximKachkov/ToDooo/internal/core/transport/http/response"
)

type CreateUserRequest struct {
	FullName    string  `json:"full_name" validate:"required,min=3,max=100"                   example:"Иван Иванов"`
	PhoneNumber *string `json:"phone_number" validate:"omitempty,min=10,max=15,startswith=+"  example:"+77081540727"`
}

type CreateUserResponse UserDTOResponse

// CreateUser 		godoc
// @Summary 		Создать пользователя
// @Description 	Создать нового пользователя в системе
// @Tags 			users
// @Accept 			json
// @Produce 		json
// @Param  			request body CreateUserRequest true "CreateUser тело запроса"
// @Success 		201 {object} CreateUserResponse "Успешно созданный пользователь"
// @Failure 		400 {object} core_http_response.ErrorResponse  "Bad request"
// @Failure 		500 {object} core_http_response.ErrorResponse  "Internal server error"
// @Router 			/users [post]
func (h *UsersHTTPHandler) CreateUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(logger, rw)

	logger.Debug("invoke create user handler")

	var request CreateUserRequest

	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "error while decoding")

		return
	}

	userDomain := domainFromDTO(request)

	userDomain, err := h.userService.CreateUser(ctx, userDomain)
	if err != nil {
		responseHandler.ErrorResponse(err, "problem creating domain user ")
	}

	response := CreateUserResponse(userDTOfromDomain(userDomain))

	responseHandler.JSONResponse(response, http.StatusCreated)
}

func domainFromDTO(dto CreateUserRequest) domain.User {
	return domain.NewUserUninitialized(dto.FullName, dto.PhoneNumber)
}
