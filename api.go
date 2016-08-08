package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/elazarl/go-bindata-assetfs"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// PontusApiContext Execution context of REST API
type ApiContext struct {
	store AppStore
}

type AppState struct {
	Name        string      `json:"name"`
	Metadata    AppMetadata `json:"metadata"`
	IsInstalled bool        `json:"isInstalled"`
}

// Serve Start server and register HTTP handlers on given address
func (ctx *ApiContext) Serve(addr string) error {
	gui := http.FileServer(&assetfs.AssetFS{
		Asset: Asset, AssetDir: AssetDir,
		AssetInfo: AssetInfo, Prefix: "gui",
	})

	router := mux.NewRouter().StrictSlash(true)
	router.Path("/api/v1/apps").Methods("GET").HandlerFunc(ctx.apps)
	router.PathPrefix("/").Handler(http.StripPrefix("/", gui))

	loggedRouter := handlers.LoggingHandler(os.Stdout, router)

	return http.ListenAndServe(addr, loggedRouter)
}

// Return all current deployments
func (ctx *ApiContext) apps(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	apps := make([]AppState, 0)
	for _, app := range ctx.store.Apps() {
		apps = append(apps, AppState{
			Name:        app.Name,
			Metadata:    app.Metadata,
			IsInstalled: ctx.store.IsInstalled(app.Name),
		})
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(apps)
}
