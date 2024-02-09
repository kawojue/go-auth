package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/kawojue/go-auth/helpers"
	"github.com/kawojue/go-auth/structs"
	"github.com/kawojue/go-hexbyte"
)

func HandleFile(
	ctx *gin.Context,
	maxSize int64,
	header *multipart.FileHeader,
	file multipart.File,
	allowedExtensions ...string,
) (*structs.File, error) {
	isAllowedExt := false
	fmt.Println(header)
	if maxSize < header.Size {
		helpers.SendError(ctx, http.StatusRequestEntityTooLarge, fmt.Sprintf("%s too large", header.Filename))
		return nil, fmt.Errorf("%s too large", header.Filename)
	}

	fileExtension := filepath.Ext(header.Filename)

	for _, ext := range allowedExtensions {
		if ext == fileExtension {
			isAllowedExt = true
			break
		}
	}

	if !isAllowedExt {
		helpers.SendError(ctx, http.StatusBadRequest, "File extension is not allowed")
		return nil, fmt.Errorf("file extension is not allowed")
	}

	fileName := hexbyte.GenerateRandomHexString(8)
	fileBytes, err := io.ReadAll(file)

	if err != nil {
		helpers.SendError(ctx, http.StatusInternalServerError, "Error reading file")
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	return &structs.File{
		FileBytes:     fileBytes,
		FileName:      fileName + "." + fileExtension,
		FileExtension: fileExtension,
	}, nil
}
