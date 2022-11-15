package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func SetupTestServer() Server {
	return Server{
		Router: gin.Default(),
	}
}

func TestPingRoute(t *testing.T) {
	testServer := SetupTestServer()
	router := testServer.setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func TestMarketRoute(t *testing.T) {
	testServer := SetupTestServer()
	router := testServer.setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/market", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestLookUpWithCollectionIdRoute_Happy(t *testing.T) {
	testServer := SetupTestServer()
	router := testServer.setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/collection/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestLookUpWithCollectionIdRoute_Error(t *testing.T) {
	testServer := SetupTestServer()
	router := testServer.setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/collection/abc", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
