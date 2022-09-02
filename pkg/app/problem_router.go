package app

import (
	"github.com/gorilla/mux"
	"github.com/kosha/freshservice-connector/pkg/httpclient"
	"net/http"
)

// getProblems godoc
// @Summary Get problems detail
// @Description List all problems
// @Tags problems
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Problems
// @Router /api/v2/problems [get]
func (a *App) getProblems(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	problems := httpclient.GetProblems(a.Cfg.GetFreshServiceURL(), a.Cfg.GetApiKey())

	respondWithJSON(w, http.StatusOK, problems)
}

// getSingleProblem godoc
// @Summary Get single problems detail
// @Description List problem by id
// @Tags problems
// @Accept  json
// @Produce  json
// @Param id path string true "Enter id e.g., 1"
// @Success 200 {object} models.SingleProblem
// @Router /api/v2/problems/{id} [get]
func (a *App) getSingleProblem(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	vars := mux.Vars(r)
	id := vars["id"]

	problem := httpclient.GetSingleProblem(a.Cfg.GetFreshServiceURL(), id, a.Cfg.GetApiKey())

	respondWithJSON(w, http.StatusOK, problem)
}

// deleteSingleProblem godoc
// @Summary Delete single problems detail
// @Description Delete problem by id
// @Tags problems
// @Accept  json
// @Produce  json
// @Param id path string true "Enter id e.g., 1"
// @Success 200 {object} string
// @Router /api/v2/problems/{id} [delete]
func (a *App) deleteSingleProblem(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	vars := mux.Vars(r)
	id := vars["id"]

	resp := httpclient.DeleteSingleProblem(a.Cfg.GetFreshServiceURL(), id, a.Cfg.GetApiKey())

	respondWithJSON(w, http.StatusOK, resp)
}
