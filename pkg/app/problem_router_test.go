package app

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func (a *App) TestGetAllProblems(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v2/problems", nil)
	checkError(err, t)
	rr := httptest.NewRecorder()
	http.HandlerFunc(a.getProblems).
		ServeHTTP(rr, req)

	//Confirm the response has the right status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d .\n Got %d instead", http.StatusOK, status)
	}
	assert.Equal(t, http.StatusOK, rr.Code)
}

func (a *App) TestGetSingleProblem(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v2/problems/1234", nil)
	checkError(err, t)
	rr := httptest.NewRecorder()
	http.HandlerFunc(a.getSingleProblem).
		ServeHTTP(rr, req)
	//Confirm the response has the right status code
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("Status code differs. Expected %d .\n Got %d instead", http.StatusOK, status)
	}
	assert.Equal(t, http.StatusOK, rr.Code)
}

