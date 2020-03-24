package script

import (
	"github.com/quicky-dev/hub-api/db"
)

type scriptError struct {
	ErrorMsg string `json:"errorMessage"`
}

type getScriptResponse struct {
	Script db.Script `json:"script"`
}
