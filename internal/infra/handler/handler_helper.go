package handler

import (
	"log/slog"
	"net/http"
)

func HandleRequestError(w http.ResponseWriter, status int, err error, detailedErr string) {
	if len(detailedErr) > 0 {
		slog.Error(detailedErr)
	}
	slog.Error("Error: " + err.Error())
	http.Error(w, err.Error(), status)
}
