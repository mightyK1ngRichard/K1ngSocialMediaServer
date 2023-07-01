package utils

import (
	"math/rand"
	"mime/multipart"
	"strconv"
	"time"
)

func GenerateUniqueFileName(expansion string) string {
	timestamp := time.Now().UnixNano()
	randomNum := rand.Intn(1000)
	fileName := strconv.FormatInt(timestamp, 10) + "_" + strconv.Itoa(randomNum) + expansion
	return fileName
}

func GetFileType(file *multipart.FileHeader) string {
	contentType := file.Header.Get("Content-Type")
	switch contentType {
	case "image/jpeg", "image/pjpeg", "image/png", "image/bmp", "image/svg+xml", "image/jpg":
		return "image"
	case "video/mp4", "video/quicktime", "video/mpeg", "video/webm", "video/x-msvideo", "video/x-flv":
		return "video"
	case "image/gif":
		return "gif"
	case "multipart/form-data":
		return "files"
	default:
		return ""
	}
}
