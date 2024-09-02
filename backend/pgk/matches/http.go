package matches

import (
	"github.com/gorilla/mux"
	"github.com/optimuscrime/hltv-no-spoilers/pgk/logger"
	"github.com/optimuscrime/hltv-no-spoilers/pgk/render"
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
		log.Debug("failed to find matches for team: %v", err)
		return
	}

	render.JSON(w, r, &resp)
}
