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

func TestGetAgents(t *testing.T) {
	// build our response JSON
	expectedserverresponse := `{
  "agents": [
    {
      "active": true,
      "auto_assign_tickets": true,
      "created_at": "2022-10-19T02:53:15Z",
      "custom_fields": {},
      "email": "mailmeatjack@proton.me",
      "first_name": "Jack",
      "has_logged_in": true,
      "id": 27002251941,
      "language": "en",
      "last_active_at": "2022-10-28T15:13:36Z",
      "last_login_at": "2022-10-19T02:53:24Z",
      "last_name": "S",
      "role_ids": [
        27000539778
      ],
      "roles": [
        {
          "role_id": 27000539778,
          "assignment_scope": "entire_helpdesk"
        }
      ],
      "scopes": {},
      "time_format": "12h",
      "time_zone": "Casablanca",
      "updated_at": "2022-10-19T02:53:15Z"
    }]}`
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
	groups := httpclient.GetAgents(server.URL+"/api/v2/agents", "foobar-test-api-key")
	bytes, err := json.Marshal(groups)
	var getAgents []models.Agents
	json.Unmarshal(bytes, &getAgents)
	assert.Equal(t, nil, err)
	if bytes == nil {
		t.Error("GetGroups API Test Fail")
	}
	t.Log("Test Execution of Get Agents Completed")
}
