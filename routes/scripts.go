package routes

import (
	"github.com/labstack/echo"
	"github.com/quicky-dev/hub-api/controllers"
)

// returns file when user runs setup script from terminal
func GetUbuntuFile(app *echo.Echo) {
	app.GET("/api/v1/os/ubuntu/scripts/:uuid", controllers.GetUbuntuFile)
}
