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

func TestGetTickets(t *testing.T) {
	// build our response JSON
	expectedserverresponse := ``
	// create a new server with that JSON
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("content-type:", "application/json")
		w.Write([]byte(expectedserverresponse))
	}))
	defer server.Close()
	groups := httpclient.GetSingleTicket(server.URL+"api/v2/tickets", "1", "foobar-api-key")
	bytes, err := json.Marshal(groups)
	var getTickets []models.Tickets
	json.Unmarshal(bytes, &getTickets)
	assert.Equal(t, nil, err)
	t.Log("Test Execution of Get Tickets Completed")
}
