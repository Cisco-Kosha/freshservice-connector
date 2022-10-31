package app

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func (a *App) TestGetAgents(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v2/agents", nil)
	checkError(err, t)
	rr := httptest.NewRecorder()
	http.HandlerFunc(a.getAgents).
		ServeHTTP(rr, req)
	//Confirm the response has the right status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d .\n Got %d instead", http.StatusOK, status)
	}
	assert.Equal(t, http.StatusOK, rr.Code)
}

func (a *App) TestGetAgentsWithID(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v2/agents/1234", nil)
	checkError(err, t)
	rr := httptest.NewRecorder()
	http.HandlerFunc(a.getSingleAgent).
		ServeHTTP(rr, req)
	//Confirm the response has the right status code
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("Status code differs. Expected %d .\n Got %d instead", http.StatusOK, status)
	}
	assert.Equal(t, http.StatusOK, rr.Code)
}



