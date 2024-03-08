package main

import (
	"os"
	"train/internal/api"
	"train/internal/domain/MainProvider"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Can't read env")
	}
	go MainProvider.StartTelegramBot()
	e := echo.New()
	anekdotProvider := MainProvider.New(
		os.Getenv("ANEKDOT_URL"),
	)
	e.GET(api.GetAnekdotHandlerName, anekdotProvider.GetAnekdot)

	port := os.Getenv("PORT")
	if port == "" {
		port = "1323"
	}
	_ = e.Start(":" + port)

}
