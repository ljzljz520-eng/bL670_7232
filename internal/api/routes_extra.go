package api

import (
	"net/http"
	"strings"
)

func routeName(path string) string { return strings.Trim(strings.TrimPrefix(path, "/"), "/") }

func isCollection(path, name string) bool { return routeName(path) == name }

func isMember(path, name string) bool {
	value := routeName(path)
	return strings.HasPrefix(value, name+"/")
}

func statusForError(err error) int {
	if err == nil {
		return http.StatusOK
	}
	if strings.Contains(err.Error(), "not found") {
		return http.StatusNotFound
	}
	if strings.Contains(err.Error(), "required") || strings.Contains(err.Error(), "invalid") {
		return http.StatusUnprocessableEntity
	}
	return http.StatusInternalServerError
}

func methodAllowed(r *http.Request, methods ...string) bool {
	for _, method := range methods {
		if r.Method == method {
			return true
		}
	}
	return false
}

func locationHeader(w http.ResponseWriter, path string) {
	if path != "" {
		w.Header().Set("Location", path)
	}
}
