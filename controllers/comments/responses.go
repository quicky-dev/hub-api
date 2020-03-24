package comments

import (
	"github.com/quicky-dev/hub-api/db"
)

type getCommentResponse struct {
	Comment db.Comment `json:"comment"`
}
