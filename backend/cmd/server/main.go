package main

import (
	"github.com/gorilla/mux"
	"github.com/optimuscrime/hltv-no-spoilers/pgk/match"
	"github.com/optimuscrime/hltv-no-spoilers/pgk/middleware"
	"github.com/optimuscrime/hltv-no-spoilers/pgk/search"
	"github.com/optimuscrime/hltv-no-spoilers/pgk/team"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	sLogger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	slog.SetDefault(sLogger)

	sLogger.Debug("Boot complete")

	r := mux.NewRouter()
	r.Use(middleware.CreateCorsMiddleware)
	r.Use(middleware.CreateLoggerMiddleware(sLogger))

	search.RegisterHandlers(r)
	team.RegisterHandlers(r)
	match.RegisterHandlers(r)

	sLogger.Debug("Starting server on port 8182")
	http.ListenAndServe(":8182", r)
}
