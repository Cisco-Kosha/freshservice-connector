package app

import (
	"encoding/json"
	"github.com/kosha/freshservice-connector/pkg/httpclient"
	"github.com/kosha/freshservice-connector/pkg/models"
	"net/http"
)

// listConnectorSpecification godoc
// @Summary Get details of the connector specification
// @Description Get all environment variables that need to be supplied
// @Tags specification
// @Accept  json
// @Produce  json
// @Success 200 {object} object
// @Router /api/v2/specification/list [get]
func (a *App) listConnectorSpecification(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	respondWithJSON(w, http.StatusOK, map[string]string{
		"API_KEY":     "Freshservice API Key",
		"DOMAIN_NAME": "Freshservice Domain Name",
	})
}

// testConnectorSpecification godoc
// @Summary Test if API key and domain name work against the specification
// @Description Check if domain account can be verified
// @Tags specification
// @Accept  json
// @Produce  json
// @Param text body models.Specification false "Enter api key and domain name properties"
// @Success 200 {object} models.Groups
// @Router /api/v2/specification/test [post]
func (a *App) testConnectorSpecification(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	if (*r).Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}

	var s models.Specification
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&s); err != nil {
		a.Log.Errorf("Error parsing json payload", err)
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	groups := httpclient.GetGroups(s.DomainName, s.ApiKey)
	if groups != nil {
		respondWithJSON(w, http.StatusOK, groups)
	} else {
		respondWithError(w, http.StatusBadRequest, "Account not verified")
	}
}
