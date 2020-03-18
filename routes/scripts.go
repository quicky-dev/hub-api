package routes

import (
	"github.com/labstack/echo"
	"github.com/quicky-dev/hub-api/controllers"
)

/* ----------------------------- Linux Packages ----------------------------- */

func GetUbuntuGeneric(app *echo.Echo) {
	app.GET("/api/v1/os/ubuntu/generic", controllers.GetUbuntuGeneric)
}

// returns object of supported downloads
func GetUbuntuItems(app *echo.Echo) {
	app.GET("/api/v1/os/ubuntu/availableItems", controllers.GetUbuntuItems)
}

// send arr's of software to setup, returns a custom setup script
func GetUbuntuCustom(app *echo.Echo) {
	app.POST("/api/v1/os/ubuntu/dynamic", controllers.GetUbuntuCustom)
}

// returns file when user runs setup script from terminal
func GetUbuntuFile(app *echo.Echo) {
	app.GET("/api/v1/os/ubuntu/scripts/:uuid", controllers.GetUbuntuFile)
}

/* ----------------------------- MacOS Packages ------------------------- */

func GetMacOSGeneric(app *echo.Echo) {
	app.GET("/api/v1/os/macos/generic", controllers.GetMacOSGeneric)
}

// returns object of supported downloads
func GetMacOSItems(app *echo.Echo) {
	app.GET("/api/v1/os/macos/availableItems", controllers.GetMacOSItems)
}

// send arr's of software to setup, returns a custom setup script
func GetMacOSCustom(app *echo.Echo) {
	app.POST("/api/v1/osmacosu/dynamic", controllers.GetMacOSCustom)
}

// returns file when user runs setup script from terminal
func GetMacOSFile(app *echo.Echo) {
	app.GET("/api/v1/os/macos/scripts/:uuid", controllers.GetMacOSFile)
}
