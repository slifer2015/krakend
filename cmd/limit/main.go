package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/api/v1/:user/:action", func(c echo.Context) error {
		fmt.Println("user", c.Param("user"))
		fmt.Println("action", c.Param("action"))
		// TODO : check and validate if user can perform this action
		return c.JSON(http.StatusOK, struct {
			Status string `json:"status"`
		}{
			Status: "OK",
		})
	})

	e.POST("/api/v1/:user/:action/apply", func(c echo.Context) error {
		fmt.Println("apply action = " + c.Param("action") + " for user = " + c.Param("user"))
		return c.JSON(http.StatusOK, nil)
	})

	e.Logger.Fatal(e.Start(":1325"))
}
