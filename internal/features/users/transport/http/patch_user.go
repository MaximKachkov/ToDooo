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
	core_http_utils "github.com/MaximKachkov/ToDooo/internal/core/transport/http/utils"
)

type PatchUserRequest struct {
	FullName    core_http_types.Nullable[string] `json:"full_name"`
	PhoneNumber core_http_types.Nullable[string] `json:"phone_number"`
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

func (s *UsersHTTPHandler) PatchUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(logger, rw)

	userId, err := core_http_utils.GetInPathValue(r, "id")
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
	var domainPatch domain.UserPatch

	domainPatch.FullName = request.FullName.ToDomainNullable()
	domainPatch.PhoneNumber = request.PhoneNumber.ToDomainNullable()

	return domainPatch
}
