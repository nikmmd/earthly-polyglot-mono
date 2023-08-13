package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	server = true
)

func main() {
	e := echo.New()
	e.GET("/one/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, fmt.Sprintf("Hello from server=%t", server))
	})
	_ = e.Start(":8080")
}
