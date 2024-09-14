package team

import (
	"github.com/gorilla/mux"
	"github.com/optimuscrime/hltv-no-spoilers/pgk/logger"
	"github.com/optimuscrime/hltv-no-spoilers/pgk/render"
	"github.com/optimuscrime/hltv-no-spoilers/pgk/resterr"
	"net/http"
)

func RegisterHandlers(r *mux.Router) {
	h := &httpHandler{}

	r.HandleFunc("/v1/team/{teamId}/matches", h.getMatchesForTeam).Methods(http.MethodGet, http.MethodOptions)
}

type httpHandler struct {
}

func (h *httpHandler) getMatchesForTeam(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log := logger.FromContext(ctx)

	teamId := mux.Vars(r)["teamId"]

	resp, err := findMatchesForTeam(teamId)
	if err != nil {
		render.JSON(w, r, resterr.FromErr(err, 500))
		return
	}

	if resp == nil {
		log.Debug("could not find any matches for team")
		w.WriteHeader(404)
		return
	}

	log.Debug("successfully fetched matches for team")

	render.JSON(w, r, resp)
}
