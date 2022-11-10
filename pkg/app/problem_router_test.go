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


func TestGetProblems(t *testing.T) {
	// build our response JSON
	expectedserverresponse := `{
  "problems": [
    {
      "id": 1,
      "requester_id": 27002251944,
      "description": "<div>Hi guys, <br/><br/>We have been facing issues when we try to reach Email Server 3. Looks like there is something wrong here.<br/><br/>Regards<br/> Rachel<br/> </div> ",
      "due_by": "2022-11-02T02:53:43Z",
      "subject": "Unable to reach email server",
      "priority": 1,
      "impact": 1,
      "status": 1,
      "created_at": "2022-10-19T02:53:43Z",
      "updated_at": "2022-10-19T02:53:43Z"
    }
  ]
}`
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
	groups := httpclient.GetProblems(server.URL + "/api/v2/groups", "foobar-test-api-key")
	bytes, err := json.Marshal(groups)
	var getProblems []models.Problems
	json.Unmarshal(bytes, &getProblems)
	assert.Equal(t, nil, err)
	if bytes == nil{
		t.Error("GetGroups API Test Fail")
	}
	t.Log("Test Execution of Get Problems Completed")
}