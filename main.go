package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const (
	USER_ACCESS = "User-Access"
	VALID_USER  = "superuser"
)

func evaluateExpression(expr string) (int, error) {
	tokens := strings.Split(expr, "+")
	result := 0

	for _, token := range tokens {
		if strings.Contains(token, "-") {
			subTokens := strings.Split(token, "-")
			subResult, err := strconv.Atoi(subTokens[0])
			if err != nil {
				return 0, err
			}

			for _, subToken := range subTokens[1:] {
				num, err := strconv.Atoi(subToken)
				if err != nil {
					return 0, err
				}
				subResult -= num
			}

			result += subResult
		} else {
			num, err := strconv.Atoi(token)
			if err != nil {
				return 0, err
			}
			result += num
		}
	}

	return result, nil
}

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

	result, err := evaluateExpression(string(body))
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
