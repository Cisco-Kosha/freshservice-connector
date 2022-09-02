package app

import (
	"github.com/gorilla/mux"
	"github.com/kosha/freshservice-connector/pkg/httpclient"
	"net/http"
)

// getGroups godoc
// @Summary Get group detail
// @Description List groups
// @Tags groups
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Groups
// @Router /api/v2/groups [get]
func (a *App) getGroups(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	groups := httpclient.GetGroups(a.Cfg.GetFreshServiceURL(), a.Cfg.GetApiKey())

	respondWithJSON(w, http.StatusOK, groups)
}

// getSingleGroup godoc
// @Summary Get single group detail
// @Description List group detail
// @Tags groups
// @Accept  json
// @Produce  json
// @Param id path string true "Enter id e.g., 1"
// @Success 200 {object} models.SingleGroup
// @Router /api/v2/groups/{id} [get]
func (a *App) getSingleGroup(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	vars := mux.Vars(r)
	id := vars["id"]

	groups := httpclient.GetSingleGroup(a.Cfg.GetFreshServiceURL(), id, a.Cfg.GetApiKey())

	respondWithJSON(w, http.StatusOK, groups)
}
