package storage

import (
	"errors"
	"fmt"
	"mime/multipart"
	"os"

	"github.com/gofiber/fiber/v2"
)

type StorageType int

const (
	Public StorageType = iota
	Private
)

func (s StorageType) String() string {
	switch s {
	case Public:
		return "storage/public/"
	case Private:
		return "storage/private/"
	default:
		return "storage/public/"
	}
}

type StorageMimeType struct {
	MimeType []string
}

func CreateStorageDirectory(storageType StorageType, directoryName string) (string, error) {
	path := fmt.Sprintf("%s%s/", storageType.String(), directoryName)
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			return "", err
		}
	}

	return path, nil
}

func Url(ctx *fiber.Ctx, path string) string {
	return ctx.BaseURL() + "/api/" + path
}

func CheckMimeType(mimeTypes []string, fileType string) bool {
	for _, mimeType := range mimeTypes {
		if fileType == mimeType {
			return true
		}
	}

	return false
}

func UploadSingle(ctx *fiber.Ctx, storageType StorageType, file *multipart.FileHeader, mimeType []string, directoryName string) (string, error) {
	fileType := file.Header.Get("Content-Type")

	mimeTypeValid := CheckMimeType(mimeType, fileType)
	if !mimeTypeValid {
		return "", errors.New(fmt.Sprintf("The file must be of the type %v", mimeType))
	}

	path, err := CreateStorageDirectory(storageType, directoryName)
	if err != nil {
		return "", errors.New("Failed to create storage directory")
	}

	filePath := path + file.Filename
	if err := ctx.SaveFile(file, filePath); err != nil {
		return "", errors.New("Failed to upload")
	}

	return filePath, nil
}
