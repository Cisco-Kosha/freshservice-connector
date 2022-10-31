package app

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)



func (a *App) TestGetGroups(t *testing.T) {

	req, err := http.NewRequest("GET", "/api/v2/groups", nil)
	checkError(err, t)
	rr := httptest.NewRecorder()
	http.HandlerFunc(a.listConnectorSpecification).
		ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d .\n Got %d instead", http.StatusOK, status)
	}
	assert.Equal(t, http.StatusOK, rr.Code)
}

