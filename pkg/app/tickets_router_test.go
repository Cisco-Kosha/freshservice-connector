package app

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func (a *App) TestGetAllTickets(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v2/tickets", nil)
	checkError(err, t)
	rr := httptest.NewRecorder()
	http.HandlerFunc(a.getAllTickets).
		ServeHTTP(rr, req)
	//Confirm the response has the right status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d .\n Got %d instead", http.StatusOK, status)
	}
	assert.Equal(t, http.StatusOK, rr.Code)
}


func (a *App) TestGetAllTicketsMetadata(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v2/tickets/metadata", nil)
	checkError(err, t)
	rr := httptest.NewRecorder()
	http.HandlerFunc(a.getAllTicketsMetadata).
		ServeHTTP(rr, req)

	//Confirm the response has the right status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d .\n Got %d instead", http.StatusOK, status)
	}
	assert.Equal(t, http.StatusOK, rr.Code)
}
