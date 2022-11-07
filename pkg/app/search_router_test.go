package app

import (
"encoding/json"
"github.com/kosha/freshservice-connector/pkg/httpclient"
"github.com/kosha/freshservice-connector/pkg/models"
"github.com/stretchr/testify/assert"
"net/http"
"net/http/httptest"
"testing"

)

func TestSearchResults(t *testing.T) {
	// build our response JSON
	expectedserverresponse := `{}`
	// create a new server with that JSON
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("content-type:", "application/json")
		w.Write([]byte(expectedserverresponse))
	}))
	defer server.Close()
	groups := httpclient.SearchTickets(server.URL+"api/v2/search", "foobar-test-api-key", "12345", "12")
	bytes, err := json.Marshal(groups)
	var getSearchResults []models.SearchResults
	json.Unmarshal(bytes, &getSearchResults)
	assert.Equal(t, nil, err)
	t.Log("Test Execution of Get Search Completed")
}
