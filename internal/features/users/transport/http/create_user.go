package users_transport_http

import (
	"net/http"

	"github.com/maksimablavatskii/Golang-ToDoap/internal/core/domain"
	core_logger "github.com/maksimablavatskii/Golang-ToDoap/internal/core/logger"
	core_http_request "github.com/maksimablavatskii/Golang-ToDoap/internal/core/transport/http/request"
	core_http_response "github.com/maksimablavatskii/Golang-ToDoap/internal/core/transport/http/response"
)

type CreateUserRequest struct {
	FullName    string  `json:"full_name" validate:"required,min=3,max=100"`
	PhoneNUmber *string `json:"phone_number" validate:"omitempty,min=10,max=15,startswith=+"`
}

type CreateUserResponse UserDTOResponse

func (h *UsersHTTPHandler) CreateUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	log.Debug("invoke CreateUser handler")
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	var request CreateUserRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request")

		return
	}

	userDomain := domainFromDTO(request)

	userDomain, err := h.usersService.CreateUser(ctx, userDomain)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create user")

		return
	}

	response := CreateUserResponse(userDTOFromDomain(userDomain))

	responseHandler.JSONResponse(response, http.StatusCreated)
}

func domainFromDTO(dto CreateUserRequest) domain.User {
	return domain.NewUserUninitialized(dto.FullName, dto.PhoneNUmber)
}
