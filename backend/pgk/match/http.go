package match

import (
	"github.com/gorilla/mux"
	"github.com/optimuscrime/hltv-no-spoilers/pgk/logger"
	"github.com/optimuscrime/hltv-no-spoilers/pgk/render"
	"github.com/optimuscrime/hltv-no-spoilers/pgk/resterr"
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
		render.JSON(w, r, resterr.FromErr(err, 500))
		return
	}

	if resp == nil {
		log.Debug("could not find any VODs for match")
		w.WriteHeader(404)
		return
	}

	log.Debug("successfully fetched match VODs")

	render.JSON(w, r, resp)
}
