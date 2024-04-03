package http

import (
	"errors"
	"net/http"
	"stori/internal/core/domain"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type errorResponse struct {
	Success  bool     `json:"success"`
	Messages []string `json:"messages"`
}

var errorStatusMap = map[error]int{}

func validationError(ctx *gin.Context, err error) {
	message := parseError(err)
	response := newErrorResponse(message)
	ctx.JSON(http.StatusBadRequest, response)
}

func parseError(err error) []string {
	var messages []string

	if errors.As(err, &validator.ValidationErrors{}) {
		for _, err := range err.(validator.ValidationErrors) {
			messages = append(messages, err.Error())
		}
	} else {
		messages = append(messages, err.Error())
	}

	return messages
}

func newErrorResponse(messages []string) errorResponse {
	return errorResponse{
		Success:  false,
		Messages: messages,
	}
}

func handleError(ctx *gin.Context, err error) {
	statusCode, ok := errorStatusMap[err]
	if !ok {
		statusCode = http.StatusInternalServerError
	}

	message := parseError(err)
	response := newErrorResponse(message)
	ctx.JSON(statusCode, response)
}

type userResponse struct {
	ID        string    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func newUserResponse(user *domain.User) userResponse {
	return userResponse{
		ID:        user.ID.String(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

type accountResponse struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Bank      string    `json:"bank"`
	Number    string    `json:"number"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func newAccountResponse(account *domain.Account) accountResponse {
	return accountResponse{
		ID:        account.ID.String(),
		UserID:    account.UserID.String(),
		Bank:      account.Bank,
		Number:    account.Number,
		CreatedAt: account.CreatedAt,
		UpdatedAt: account.UpdatedAt,
	}
}

type transactionResponse struct {
	ID        string    `json:"id"`
	AccountID string    `json:"account_id"`
	Date      string    `json:"date"`
	Credit    bool      `json:"credit"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func newTransactionResponse(transaction *domain.Transaction) transactionResponse {
	return transactionResponse{
		ID:        transaction.ID.String(),
		AccountID: transaction.AccountID.String(),
		Date:      transaction.Date.String(),
		Credit:    transaction.Credit,
		Amount:    transaction.Amount.InexactFloat64(),
		CreatedAt: transaction.CreatedAt,
		UpdatedAt: transaction.UpdatedAt,
	}
}

func handleSuccess(ctx *gin.Context, data any) {
	rsp := newResponse(true, "success", data)
	ctx.JSON(http.StatusOK, rsp)
}

type response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func newResponse(success bool, message string, data any) response {
	return response{
		Success: success,
		Message: message,
		Data:    data,
	}
}

func newTransactionsResponse(transactionsDomain []domain.Transaction) []transactionResponse {
	transactions := make([]transactionResponse, len(transactionsDomain))

	for i, transaction := range transactionsDomain {
		transactions[i] = newTransactionResponse(&transaction)
	}

	return transactions
}
