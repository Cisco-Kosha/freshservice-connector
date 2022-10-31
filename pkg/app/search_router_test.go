package main

import (
	"github.com/Cisco-Kosha/freshservice-connector/pkg/app"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetRequest(t *testing.T) {
	t.Parallel()
	r, _ := http.NewRequest("GET", "/api/v2/agents", nil)
	w := httptest.NewRecorder()
	a := app.App{}
	a.searchTickets(w,r)
	assert.Equal(t, http.StatusOK, w.Code)
}
