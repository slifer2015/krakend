package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.POST("/api/v1/protected-api/:user_id", func(c echo.Context) error {
		fmt.Println("calling protected-api")
		fmt.Println(c.Request().Header)
		gg, _ := ioutil.ReadAll(c.Request().Body)
		fmt.Println(string(gg))
		userID := c.Param("user_id")
		if userID == "" {
			return c.JSON(http.StatusUnauthorized, nil)
		}

		return c.JSON(http.StatusOK, struct {
			Status string `json:"status"`
		}{
			Status: "OK",
		})
	})
	e.Logger.Fatal(e.Start(":1324"))
}
