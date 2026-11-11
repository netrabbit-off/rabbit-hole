package middleware

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

type Response struct {
	UserID string `json:"message"`
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header["Authorization"]
		if len(header) == 0 {
			http.Error(w, "UNAUTHORIZED", http.StatusUnauthorized)
			return
		}
		token := strings.TrimPrefix(header[0], "Bearer ")
		if len(token) == 0 {
			http.Error(w, "UNAUTHORIZED", http.StatusUnauthorized)
			return
		}
		println("token: " + token)

		client := &http.Client{}
		req, err := http.NewRequest("GET", "http://127.0.0.1:8089/check", nil)
		if err != nil {
			http.Error(w, "UNAUTHORIZED", http.StatusUnauthorized)
			return
		}
		req.Header.Add("Authorization", "Bearer "+token)
		res, err := client.Do(req)
		if err != nil {
			http.Error(w, "UNAUTHORIZED", http.StatusUnauthorized)
			return
		}
		if res.Status != "200 OK" {
			http.Error(w, "UNAUTHORIZED", http.StatusUnauthorized)
			return
		}

		resp := &Response{}
		body, err := io.ReadAll(res.Body)
		if err != nil {
			http.Error(w, "UNAUTHORIZED", http.StatusUnauthorized)
			return
		}
		if err := json.Unmarshal(body, &resp); err != nil {
			http.Error(w, "UNAUTHORIZED", http.StatusUnauthorized)
			return
		}
		userID := resp.UserID
		println("user_id: " + userID)

		next.ServeHTTP(w, r)
	})
}
