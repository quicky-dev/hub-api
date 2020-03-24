package main

import (
	"os"

	"github.com/labstack/echo"
	"github.com/quicky-dev/hub-api/routers"
)

func main() {
	app := echo.New()

	routers.RegisterScriptsRoutes(app)

	app.Logger.Fatal(app.Start(os.Getenv("PORT")))

}
