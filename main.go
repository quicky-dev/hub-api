package main

import (
	"os"

	"github.com/labstack/echo"
	"github.com/quicky-dev/hub-api/routers"
	"github.com/quicky-dev/hub-api/util"
)

func main() {
	app := echo.New()

	routers.RegisterScriptsRoutes(app)

	util.MongoClient.Database("Test").Collection("User")
	app.Logger.Fatal(app.Start(os.Getenv("PORT")))

}
