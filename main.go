package main

import (
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/wazsone/softweather-test/internal/services"
)

const (
	USER_ACCESS = "User-Access"
	VALID_USER  = "superuser"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return

	}
	if user := r.Header.Get(USER_ACCESS); user != VALID_USER {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Faild to read request", http.StatusBadRequest)
		return
	}

	result, err := services.EvaluateExpression(string(body))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte(strconv.Itoa(result)))
}

func main() {
	http.HandleFunc("/calc", Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
