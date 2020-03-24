package comments

import (
	"github.com/labstack/echo"
	"github.com/quicky-dev/hub-api/db"
	"net/http"
)

func (c echo.Context) GetComment() error {
	commentID := c.Param(":commentID")

	comment, err := db.GetCommentByID(commentID)

	if err != nil {
		return c.JSON(http.StatusNotFound, getCommentResponse{comment})
	}
}
