package main

import (
	"carrotAuction/db"
	"carrotAuction/img"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main(){
	e := echo.New()
	e.GET("/", func(c echo.Context) error{
		db.Connection();
		return c.String(http.StatusOK, "Hello, World!")
	})

	// 이미지 업로드 핸들러
	e.POST("/upload", img.UploadHandler)

	// 이미지 서빙 핸들러
	e.GET("/images/:filename", img.ServeImageHandler)
	e.Logger.Fatal(e.Start(":3000"))
}




