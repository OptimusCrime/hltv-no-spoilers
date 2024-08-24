package search

import (
	"github.com/gorilla/mux"
	"github.com/optimuscrime/hltv-no-spoilers/pgk/logger"
	"github.com/optimuscrime/hltv-no-spoilers/pgk/render"
	"net/http"
)

func RegisterHandlers(r *mux.Router) {
	h := &httpHandler{}

	r.HandleFunc("/v1/search", h.Search).Methods(http.MethodGet, http.MethodOptions)
}

type httpHandler struct {
}

func (h *httpHandler) Search(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)

	term := r.URL.Query().Get("term")

	if len(term) < 2 {
		log.Debug("tried to search for term with two or fewer characters, which is not allowed")
		return
	}

	resp, err := searchForTeams(term)
	if err != nil {
		log.Debug("failed to search for teams: %v", err)
		return
	}

	render.JSON(w, r, resp)
}
