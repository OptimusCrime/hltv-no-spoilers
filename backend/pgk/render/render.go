package render

import (
	"encoding/json"
	"github.com/optimuscrime/hltv-no-spoilers/pgk/logger"
	"net/http"
)

func JSON(w http.ResponseWriter, r *http.Request, v any) {
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(v); err != nil {
		fail(w, r, err)
		return
	}
}

func fail(w http.ResponseWriter, r *http.Request, err error) {
	log := logger.FromContext(r.Context())

	log.Error("Failed to parse JSON", err)

	w.WriteHeader(http.StatusInternalServerError)
}
