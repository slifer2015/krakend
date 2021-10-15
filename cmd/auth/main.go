package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/api/v1/validate-token", func(c echo.Context) error {
		fmt.Println("calling validate token")
		authHeader := c.Request().Header.Get("Authorization")

		if authHeader == "correct header" {
			return c.JSON(http.StatusOK, struct {
				UserID string `json:"user_id"`
				Action string `json:"action"`
			}{
				UserID: "sample user id",
				Action: "sign-in",
			})
		}
		return c.JSON(http.StatusUnauthorized, nil)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
