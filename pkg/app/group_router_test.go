package app

import (
	"encoding/json"
	"github.com/kosha/freshservice-connector/pkg/httpclient"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

type Groups struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func TestGetGroups(t *testing.T) {
	// build our response JSON
	expectedserverresponse := `{
	"groups": [{
      "id": 27000688460,
      "name": "Capacity Management Team",
      "description": "Capacity Management Team",
      "created_at": "2022-10-19T02:53:16Z",
      "updated_at": "2022-10-19T02:53:16Z"
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
	groups := httpclient.GetGroups(server.URL + "/api/v2/groups", "foobar-test-api-key")
	bytes, err := json.Marshal(groups)
	var getGroups []Groups
	json.Unmarshal(bytes, &getGroups)
	assert.Equal(t, nil, err)
	if bytes == nil{
		t.Error("GetGroups API Test Fail")
	}
	t.Log("Test Execution of Get Groups Completed")
}

func AssertEqual(t *testing.T, a interface{}, b interface{}) {
	if a == b {
		return
	}
	// debug.PrintStack()
	t.Errorf("Received %v (type %v), expected %v (type %v)", a, reflect.TypeOf(a), b, reflect.TypeOf(b))
}