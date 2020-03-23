package script

import (
	"github.com/labstack/echo"
	"github.com/quicky-dev/hub-api/db"
)

// GetScripts returns all ratings and comments for scripts
func GetScripts(c echo.Context) error {

	scripts, err := db.GetScripts()

	if err != nil {
		return c.JSON(400, scriptError{"Login failed."})
	}

	return scripts
}
