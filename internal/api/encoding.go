package api

import (
	"encoding/json"
	"net/http"
)

func decodeJSON(r *http.Request, target any) error {
	if r.Body == nil {
		return nil
	}
	return json.NewDecoder(r.Body).Decode(target)
}

func setCacheHeaders(w http.ResponseWriter) {
	w.Header().Set("Cache-Control", "no-store")
	w.Header().Set("X-Content-Type-Options", "nosniff")
}

func writeText(w http.ResponseWriter, status int, content string) {
	setCacheHeaders(w)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(status)
	_, _ = w.Write([]byte(content))
}

func writeCSV(w http.ResponseWriter, status int, content string) {
	setCacheHeaders(w)
	w.Header().Set("Content-Type", "text/csv; charset=utf-8")
	w.WriteHeader(status)
	_, _ = w.Write([]byte(content))
}

func methodIs(r *http.Request, method string) bool { return r.Method == method }

func acceptsJSON(r *http.Request) bool {
	value := r.Header.Get("Accept")
	return value == "" || value == "application/json" || value == "*/*"
}
