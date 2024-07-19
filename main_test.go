package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	// Create a new request to the root endpoint
	req, err := http.NewRequest("GET", "/", nil)
	assert.NoError(t, err)

	// Create a response recorder to record the response
	rr := httptest.NewRecorder()

	// Call the main function, passing in the response recorder and request
	indexHandler(rr, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, rr.Code)

	// Check the response body
	expectedBody := "welcome people to the world"
	assert.Equal(t, expectedBody, rr.Body.String())
}
func TestCreateRouter(t *testing.T) {
	r := createRouter()

	assert.NotNil(t, r)
	assert.NotNil(t, r.NotFoundHandler)
	assert.NotNil(t, r.MethodNotAllowedHandler)
}
