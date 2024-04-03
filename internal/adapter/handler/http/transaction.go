package http

import (
	"mime/multipart"
	"stori/internal/core/domain"
	"stori/internal/core/port"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

const (
	timeFormatMonthDay = "1/2"
)

type TransactionHandler struct {
	service port.TransactionService
}

func NewTransactionHandler(service port.TransactionService) *TransactionHandler {
	return &TransactionHandler{
		service,
	}
}

type createTransactionRequest struct {
	AccountID string  `json:"account_id" binding:"required"`
	Date      string  `json:"date" binding:"required"`
	Credit    *bool   `json:"credit" binding:"required"`
	Amount    float64 `json:"amount" binding:"required"`
}

func (th *TransactionHandler) CreateTransaction(ctx *gin.Context) {
	var request createTransactionRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		validationError(ctx, err)
		return
	}

	accountID, err := uuid.Parse(request.AccountID)
	if err != nil {
		validationError(ctx, err)
		return
	}

	date, err := time.Parse(timeFormatMonthDay, request.Date)
	if err != nil {
		validationError(ctx, err)
		return
	}

	transaction := domain.Transaction{
		AccountID: accountID,
		Date:      date,
		Credit:    *request.Credit,
		Amount:    decimal.NewFromFloat(request.Amount),
	}

	_, err = th.service.CreateTransaction(ctx, &transaction)
	if err != nil {
		handleError(ctx, err)
		return
	}

	response := newTransactionResponse(&transaction)

	handleSuccess(ctx, response)
}

type getTransactionsByAccountIDRequest struct {
	AccountID string `uri:"account_id" binding:"required"`
}

func (th *TransactionHandler) GetTransactionsByAccountId(ctx *gin.Context) {
	var request getTransactionsByAccountIDRequest
	if err := ctx.ShouldBindUri(&request); err != nil {
		validationError(ctx, err)
		return
	}

	accountID, err := uuid.Parse(request.AccountID)
	if err != nil {
		validationError(ctx, err)
		return
	}

	transactions, err := th.service.GetTransactionsByAccountID(ctx, accountID)
	if err != nil {
		validationError(ctx, err)
		return
	}

	response := newTransactionsResponse(transactions)

	handleSuccess(ctx, response)
}

type uploadFileRequest struct {
	AccountID string                `form:"account_id" binding:"required"`
	File      *multipart.FileHeader `form:"file" binding:"required"`
}

func (th *TransactionHandler) UploadFile(ctx *gin.Context) {
	var request uploadFileRequest
	if err := ctx.ShouldBind(&request); err != nil {
		validationError(ctx, err)
		return
	}

	accountID, err := uuid.Parse(request.AccountID)
	if err != nil {
		validationError(ctx, err)
		return
	}

	if err := th.service.UploadFile(ctx, &domain.UploadFile{
		AccountID: accountID,
		File:      request.File,
	}); err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, nil)
}

type sendEmailSummary struct {
	AccountID string `json:"account_id" binding:"required"`
}

func (th *TransactionHandler) SendEmailSummary(ctx *gin.Context) {
	var request sendEmailSummary
	if err := ctx.ShouldBindJSON(&request); err != nil {
		validationError(ctx, err)
		return
	}

	accountID, err := uuid.Parse(request.AccountID)
	if err != nil {
		validationError(ctx, err)
		return
	}

	if err := th.service.SendEmailSummary(ctx, accountID); err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, nil)
}
