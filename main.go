package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	fmt.Println("welcome to echo !")
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"hello": "world"})

	})
	e.Logger.Fatal(e.Start(":1313"))
}
