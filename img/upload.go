package img

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

func UploadHandler(c echo.Context) error {
	// form 필드 이름은 "image"로 가정합니다.
	form, err := c.MultipartForm()
	if err != nil {
		return c.String(http.StatusBadRequest, "Failed to upload image")
	}
	files := form.File["image"]

	for _, file := range files{
		src, err := file.Open()
		if err != nil {
			return c.String(http.StatusInternalServerError, "Failed to open uploaded file")
		}
		defer src.Close()

		// 업로드된 파일을 저장할 경로와 파일 이름 생성
		filename := generateFilename(file.Filename)
		dstPath := filepath.Join("./uploads", filename)

		dst, err := os.Create(dstPath)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Failed to save image")
		}
		defer dst.Close()

		// 업로드된 파일을 저장
		if _, err = io.Copy(dst, src); err != nil {
			return c.String(http.StatusInternalServerError, "Failed to save image")
		}
	}

	return c.String(http.StatusOK, "Image uploaded successfully")
}

func ServeImageHandler(c echo.Context) error {
	filename := c.Param("filename")
	imagePath := filepath.Join("./uploads", filename)

	return c.File(imagePath)
}

func generateFilename(originalFilename string) string {
	extension := filepath.Ext(originalFilename)
	// filename := strings.TrimSuffix(originalFilename, extension)

	return fmt.Sprintf("%s%s", generateRandomString(10), extension)
}

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}