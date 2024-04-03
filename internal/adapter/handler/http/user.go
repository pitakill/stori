package http

import (
	"stori/internal/core/domain"
	"stori/internal/core/port"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service port.UserService
}

func NewUserHandler(service port.UserService) *UserHandler {
	return &UserHandler{
		service,
	}
}

type registerRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
}

func (uh *UserHandler) Register(ctx *gin.Context) {
	var request registerRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		validationError(ctx, err)
		return
	}

	user := domain.User{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
	}

	_, err := uh.service.Register(ctx, &user)
	if err != nil {
		handleError(ctx, err)
		return
	}

	response := newUserResponse(&user)

	handleSuccess(ctx, response)
}

type getUserByEmailRequest struct {
	Email string `uri:"email" binding:"required,email"`
}

func (uh *UserHandler) GetUserByEmail(ctx *gin.Context) {
	var request getUserByEmailRequest
	if err := ctx.ShouldBindUri(&request); err != nil {
		validationError(ctx, err)
		return
	}

	user, err := uh.service.GetUserByEmail(ctx, request.Email)
	if err != nil {
		validationError(ctx, err)
		return
	}

	response := newUserResponse(user)

	handleSuccess(ctx, response)
}
