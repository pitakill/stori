package domain

import (
	"mime/multipart"

	"github.com/google/uuid"
)

type UploadFile struct {
	AccountID uuid.UUID
	File      *multipart.FileHeader
}
