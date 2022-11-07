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

func TestGetAssets(t *testing.T) {
	// build our response JSON
	expectedserverresponse := `{
  "assets": [
    {
      "id": 27001019809,
      "display_id": 3,
      "name": "Logitech Mouse",
      "asset_type_id": 27002275571,
      "impact": "low",
      "author_type": "User",
      "usage_type": "permanent",
      "asset_tag": "ASSET-3",
      "created_at": "2022-10-19T02:53:47Z",
      "updated_at": "2022-10-19T02:53:47Z"
    }]}
`
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
	groups := httpclient.GetAssets(server.URL+"/api/v2/agents", "foobar-test-api-key")
	bytes, err := json.Marshal(groups)
	var getAssets []models.Assets
	json.Unmarshal(bytes, &getAssets)
	assert.Equal(t, nil, err)
	if bytes == nil {
		t.Error("GetGroups API Test Fail")
	}
	t.Log("Test Execution of Get Assets Completed")
}
