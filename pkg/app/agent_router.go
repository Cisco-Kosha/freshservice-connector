package app

import (
	"github.com/gorilla/mux"
	"github.com/kosha/freshservice-connector/pkg/httpclient"
	"net/http"
)

// getAgents godoc
// @Summary Get freshservice agents
// @Description List all agents
// @Tags agents
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Agents
// @Router /api/v2/agents [get]
func (a *App) getAgents(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	agents := httpclient.GetAgents(a.Cfg.GetFreshServiceURL(), a.Cfg.GetApiKey())

	respondWithJSON(w, http.StatusOK, agents)
}

// getSingleAgent godoc
// @Summary Get freshservice agent by id
// @Description List agent by id
// @Tags agents
// @Accept  json
// @Produce  json
// @Param id path string true "Agent id"
// @Success 200 {object} models.Agent
// @Router /api/v2/agents/{id} [get]
func (a *App) getSingleAgent(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	vars := mux.Vars(r)
	id := vars["id"]

	agents := httpclient.GetSingleAgent(a.Cfg.GetFreshServiceURL(), id, a.Cfg.GetApiKey())

	respondWithJSON(w, http.StatusOK, agents)
}

// deactivateSingleAgent godoc
// @Summary Deactivate freshservice agent by id
// @Description Deactivate agent by id
// @Tags agents
// @Accept  json
// @Produce  json
// @Param id path string true "Agent id"
// @Success 200 {object} models.Agent
// @Router /api/v2/agents/{id} [delete]
func (a *App) deactivateSingleAgent(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	vars := mux.Vars(r)
	id := vars["id"]

	agents := httpclient.DeactivateSingleAgent(a.Cfg.GetFreshServiceURL(), id, a.Cfg.GetApiKey())

	respondWithJSON(w, http.StatusOK, agents)
}
