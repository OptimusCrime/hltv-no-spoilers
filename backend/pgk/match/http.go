package match

import (
	"github.com/gorilla/mux"
	"github.com/optimuscrime/hltv-no-spoilers/pgk/logger"
	"github.com/optimuscrime/hltv-no-spoilers/pgk/render"
	"net/http"
)

func RegisterHandlers(r *mux.Router) {
	h := &httpHandler{}

	r.HandleFunc("/v1/match/{matchId}", h.getMatchVODs).Methods(http.MethodGet, http.MethodOptions)
}

type httpHandler struct {
}

func (h *httpHandler) getMatchVODs(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log := logger.FromContext(ctx)

	matchId := mux.Vars(r)["matchId"]

	resp, err := findMatchVODs(matchId)
	if err != nil {
		log.Debug("failed to find VODs for match: %v", err)
		return
	}

	render.JSON(w, r, &resp)
}
