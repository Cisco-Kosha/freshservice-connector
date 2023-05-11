package app

import (
	"io/ioutil"
	"net/http"
)

//Custom function to be implemented by connector developer
//Prints data sent to this connector via /trigger API.
func (a *App) processTrigger(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, nil)
	}
	a.Log.Infof("Received %v", string(bodyBytes))

	respondWithJSON(w, http.StatusOK, nil)
}
