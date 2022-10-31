package main

import (
	"encoding/json"
	"github.com/Cisco-Kosha/freshservice-connector/pkg/app"
	"github.com/Cisco-Kosha/freshservice-connector/pkg/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRequest(t *testing.T) {
	t.Parallel()

	r, _ := http.NewRequest("GET", "/api/v2/agents", nil)
	w := httptest.NewRecorder()

	a := app.App{}
	a.ListConnectorSpecification(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}
