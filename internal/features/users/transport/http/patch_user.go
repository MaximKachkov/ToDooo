package users_transport_http

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/MaximKachkov/ToDooo/internal/core/domain"
	core_logger "github.com/MaximKachkov/ToDooo/internal/core/logger"
	core_http_request "github.com/MaximKachkov/ToDooo/internal/core/transport/http/request"
	core_http_response "github.com/MaximKachkov/ToDooo/internal/core/transport/http/response"
	core_http_types "github.com/MaximKachkov/ToDooo/internal/core/transport/http/types"
)

type PatchUserRequest struct {
	FullName    core_http_types.Nullable[string] `json:"full_name"    swaggertype:"string" example:"Maxat Maxatovich"`
	PhoneNumber core_http_types.Nullable[string] `json:"phone_number"  swaggertype:"string" example:"+77081540727"`
}

func (r *PatchUserRequest) Validate() error {
	if r.FullName.Set {
		if r.FullName.Value == nil {
			return fmt.Errorf("`FullName` cant be NULL")
		}
		fullNameLen := len([]rune(*r.FullName.Value))
		if fullNameLen < 3 || fullNameLen > 100 {
			return fmt.Errorf("not appropraite length : %d", fullNameLen)
		}
	}
	if r.PhoneNumber.Set {
		if r.PhoneNumber.Value != nil {
			phoneNumberLen := len([]rune(*r.PhoneNumber.Value))
			if phoneNumberLen < 10 || phoneNumberLen > 15 {
				return fmt.Errorf("phone number must be beteween 10 and 15 symbols")
			}
			if !strings.HasPrefix(*r.PhoneNumber.Value, "+") {
				return fmt.Errorf("phone number must start with + symbol")
			}
		}
	}

	return nil
}

type PatchUserResponse UserDTOResponse

// PatchUser    godoc
// @Summary  Patch  user
// @Description Patch and existing user by id
// @Description ### Логика обновления полей (Three-state logic):
// @Description 1. ***Поле не передано***: `phone_number` игнорируется , значение в БД не меняется
// @Description 2. ***Явно передано значение***: `"phone_number":"+77081540727"` - устанавливает новый номер в бд
// @Description 3. **Передан null**: `"phone_number": null`	 null - очищает значение в дб
// @Description 4. **Ограничения**: `full_name` не может быть выставлен как null
// @Tags  users
// @Accept json
// @Produce json
// @Param   id path int true  "id of user"
// @Param   request body PatchUserRequest true  "request body"
// @Success 200 {object} PatchUserResponse "Patched user"
// @Failure 400 {object} core_http_response.ErrorResponse  "Bad request"
// @Failure 404 {object} core_http_response.ErrorResponse  "Not Found"
// @Failure 409 {object} core_http_response.ErrorResponse  "Conflict"
// @Failure 500 {object} core_http_response.ErrorResponse  "Internal server error"
// @Router  /users/{id} [patch]
func (s *UsersHTTPHandler) PatchUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(logger, rw)

	userId, err := core_http_request.GetInPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "error getinpathvalue patchuser")
		return
	}

	var request PatchUserRequest

	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "error decoding/validating patch user request")
		return
	}

	domainPatch := UserPatchFromRequest(request)

	domainUser, err := s.userService.PatchUser(ctx, userId, domainPatch)

	if err != nil {
		responseHandler.ErrorResponse(err, "error patching user")
		return
	}
	userDTO := PatchUserResponse(userDTOfromDomain(domainUser))
	responseHandler.JSONResponse(userDTO, http.StatusOK)

}

func UserPatchFromRequest(request PatchUserRequest) domain.UserPatch {

	return domain.NewUserPatch(request.FullName.ToDomainNullable(), request.PhoneNumber.ToDomainNullable())
}
