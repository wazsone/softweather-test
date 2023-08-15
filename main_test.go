package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	tests := []struct {
		name       string
		method     string
		path       string
		user       string
		response   string
		expression []byte
		status     int
	}{
		{
			name:       "Test GET /calc",
			method:     "GET",
			path:       "/calc",
			user:       "superuser",
			expression: []byte("2+2-3-5+1"),
			response:   "Method not allowed\n",
			status:     http.StatusMethodNotAllowed,
		},
		{
			name:       "Test POST /calc",
			method:     "POST",
			path:       "/calc",
			user:       "superuser",
			expression: []byte("2+2-3-5+1"),
			response:   "-3",
			status:     http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(tt.method, tt.path, bytes.NewBuffer(tt.expression))
			if err != nil {
				t.Fatalf("failed to create request: %v", err)
			}
			req.Header.Set(USER_ACCESS, tt.user)

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(Handler)
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.status {
				t.Errorf("handler returned wrong status code: got %v, want %v",
					status, http.StatusOK)
			}

			if rr.Body.String() != tt.response {
				t.Errorf("handler returned unexpected body: got %v, want %v",
					rr.Body.String(), tt.response)
			}
		})
	}
}
