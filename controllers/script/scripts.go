package script

import (
	"github.com/labstack/echo"
	"github.com/quicky-dev/hub-api/db"
)

// GetScripts returns all ratings and comments for scripts
func GetScripts(c echo.Context) error {

	scripts, err := db.ReturnScripts()

	if err != nil {
		return c.JSON(400, scriptError{"No scripts"})
	}

	return c.JSON(200, scripts)
}

// UpdateScript updates the script obj with associated changes
func UpdateScript(c echo.Context) error {

}
