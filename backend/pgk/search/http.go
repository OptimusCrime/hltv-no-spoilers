package search

import (
	"github.com/gorilla/mux"
	"github.com/optimuscrime/hltv-no-spoilers/pgk/logger"
	"github.com/optimuscrime/hltv-no-spoilers/pgk/render"
	"github.com/optimuscrime/hltv-no-spoilers/pgk/resterr"
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
		render.JSON(w, r, resterr.New("at least two characters are required when searching for team", 400))
		return
	}

	resp, err := searchForTeams(term)
	if err != nil {
		render.JSON(w, r, resterr.FromErr(err, 500))
		return
	}

	log.Debug("successfully searched for team by term")

	render.JSON(w, r, resp)
}
