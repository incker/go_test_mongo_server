package server

import (
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIServer_HandleHello(t *testing.T) {
	config, err := NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	apiServer, err := New(config)
	if err != nil {
		log.Fatal(err)
	}

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/hello", nil)
	handleHello(apiServer).ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "Hello")
}
