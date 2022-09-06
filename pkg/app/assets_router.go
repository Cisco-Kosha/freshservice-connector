package app

import (
	"github.com/gorilla/mux"
	"github.com/kosha/freshservice-connector/pkg/httpclient"
	"net/http"
)

// getAssets godoc
// @Summary Get assets details
// @Description List all assets
// @Tags assets
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Assets
// @Failure      400  {object} string "bad request"
// @Failure      403  {object}  string "permission denied"
// @Failure      404  {object}  string "not found"
// @Failure      500  {object}  string "internal server error"
// @Router /api/v2/assets [get]
func (a *App) getAssets(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	assets := httpclient.GetAssets(a.Cfg.GetFreshServiceURL(), a.Cfg.GetApiKey())

	respondWithJSON(w, http.StatusOK, assets)
}

// getSingleAsset godoc
// @Summary Get single assets detail
// @Description List asset by id
// @Tags assets
// @Accept  json
// @Produce  json
// @Param id path string true "Enter id e.g., 1"
// @Success 200 {object} models.SingleAsset
// @Failure      400  {object} string "bad request"
// @Failure      403  {object}  string "permission denied"
// @Failure      404  {object}  string "not found"
// @Failure      500  {object}  string "internal server error"
// @Router /api/v2/assets/{id} [get]
func (a *App) getSingleAsset(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	vars := mux.Vars(r)
	id := vars["id"]

	asset := httpclient.GetSingleAsset(a.Cfg.GetFreshServiceURL(), id, a.Cfg.GetApiKey())

	respondWithJSON(w, http.StatusOK, asset)
}

// deleteSingleProblem godoc
// @Summary Delete single asset
// @Description Delete asset by id
// @Tags assets
// @Accept  json
// @Produce  json
// @Param id path string true "Enter id e.g., 1"
// @Success 200 {object} string
// @Failure      400  {object} string "bad request"
// @Failure      403  {object}  string "permission denied"
// @Failure      404  {object}  string "not found"
// @Failure      500  {object}  string "internal server error"
// @Router /api/v2/assets/{id} [delete]
func (a *App) deleteSingleAsset(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	vars := mux.Vars(r)
	id := vars["id"]

	agents := httpclient.DeleteSingleAsset(a.Cfg.GetFreshServiceURL(), id, a.Cfg.GetApiKey())

	respondWithJSON(w, http.StatusOK, agents)
}
