package http

import (
	"stori/internal/core/domain"
	"stori/internal/core/port"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AccountHandler struct {
	service port.AccountService
}

func NewAccountHandler(service port.AccountService) *AccountHandler {
	return &AccountHandler{
		service,
	}
}

type createAccountRequest struct {
	UserID string `json:"user_id" binding:"required"`
	Bank   string `json:"bank" binding:"required"`
	Number string `json:"number" binding:"required"`
}

func (ah *AccountHandler) CreateAccount(ctx *gin.Context) {
	var request createAccountRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		validationError(ctx, err)
		return
	}

	userId, err := uuid.Parse(request.UserID)
	if err != nil {
		validationError(ctx, err)
		return
	}

	account := domain.Account{
		UserID: userId,
		Bank:   request.Bank,
		Number: request.Number,
	}

	_, err = ah.service.CreateAccount(ctx, &account)
	if err != nil {
		handleError(ctx, err)
		return
	}

	response := newAccountResponse(&account)

	handleSuccess(ctx, response)
}

type getAccountByUserIDRequest struct {
	UserID string `uri:"user_id" binding:"required"`
}

func (ah *AccountHandler) GetAccountByUserID(ctx *gin.Context) {
	var request getAccountByUserIDRequest
	if err := ctx.ShouldBindUri(&request); err != nil {
		handleError(ctx, err)
		return
	}

	userId, err := uuid.Parse(request.UserID)
	if err != nil {
		validationError(ctx, err)
		return
	}

	account, err := ah.service.GetAccountByUserID(ctx, userId)
	if err != nil {
		handleError(ctx, err)
		return
	}

	response := newAccountResponse(account)

	handleSuccess(ctx, response)
}
