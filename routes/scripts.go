package routes

import (
	"github.com/labstack/echo"
	"github.com/quicky-dev/hub-api/controllers"
)

func RegisterScriptsRoutes(app *echo.Echo) {
	/* ----------------------------- Linux Packages ------------------------- */
	app.GET("/api/v1/os/ubuntu/generic", controllers.GetUbuntuGeneric)
	// returns object of supported downloads
	app.GET("/api/v1/os/ubuntu/availableItems", controllers.GetUbuntuItems)
	// send arr's of software to setup, returns a custom setup script
	app.POST("/api/v1/os/ubuntu/dynamic", controllers.GetUbuntuCustom)
	// returns file when user runs setup script from terminal
	app.GET("/api/v1/os/ubuntu/scripts/:uuid", controllers.GetUbuntuFile)

	/* ----------------------------- MacOS Packages ------------------------- */
	app.GET("/api/v1/os/macos/generic", controllers.GetMacOSGeneric)
	app.GET("/api/v1/os/macos/availableItems", controllers.GetMacOSItems)
	app.POST("/api/v1/osmacosu/dynamic", controllers.GetMacOSCustom)
	app.GET("/api/v1/os/macos/scripts/:uuid", controllers.GetMacOSFile)

}
