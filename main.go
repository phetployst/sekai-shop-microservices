package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/phetployst/sekai-shop-microservices/config"
)

func main() {
	e := echo.New()
	ctx := context.Background()
	_ = ctx

	cfg := config.LoadConfig(func() string {
		if len(os.Args) < 2 {
			log.Fatal("Error: .env path is required")
		}
		return os.Args[1]
	}())

	log.Println(cfg)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!!")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
