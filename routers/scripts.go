package routers

import (
	"github.com/labstack/echo"
	"github.com/quicky-dev/hub-api/controllers/comments"
	"github.com/quicky-dev/hub-api/controllers/script"
)

func RegisterScriptsRoutes(app *echo.Echo) {
	// Scripts
	app.GET("/api/v1/scripts", script.GetScripts)
	app.GET("/api/v1/scripts/:id", script.GetScriptID)
	app.PUT("/api/v1/scripts/:id", script.UpdateScript)

	// Comments
	app.GET("/api/v1/comments", comments.GetComment)

}
