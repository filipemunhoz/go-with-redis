package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthCheckHandler(t *testing.T) {

	req, err := http.NewRequest(http.MethodGet, "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetHealth)

	handler.ServeHTTP(rr, req)

	checkSuccessStatus(rr.Code, t)

	// Check the response body is what we expect.
	expected := `{"status": "ok"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

	assert.JSONEq(t, expected, rr.Body.String(), "JSON Response not equals what expected")

}

func checkSuccessStatus(code int, t *testing.T) {
	if code != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, code)
	}
}
