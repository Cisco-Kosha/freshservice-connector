package app

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func checkError(err error, t *testing.T) {
	if err != nil {
		t.Errorf("An error occurred. %v", err)
	}
}

func (a *App) TestSpecificationListHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v2/specification/list", nil)
	checkError(err, t)
	rr := httptest.NewRecorder()
	http.HandlerFunc(a.listConnectorSpecification).
		ServeHTTP(rr, req)
	//Confirm the response has the right status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d .\n Got %d instead", http.StatusOK, status)
	}
	assert.Equal(t, http.StatusOK, rr.Code)
}

func (a *App) TestSearchTickets(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v2/search", nil)
	checkError(err, t)
	rr := httptest.NewRecorder()
	http.HandlerFunc(a.listConnectorSpecification).
		ServeHTTP(rr, req)
	//Confirm the response has the right status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d .\n Got %d instead", http.StatusOK, status)
	}
	assert.Equal(t, http.StatusOK, rr.Code)
}
