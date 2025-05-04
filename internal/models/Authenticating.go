package main

import (
	"fmt"
	"net/http"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session_token")
	if err == http.ErrNoCookie {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusBadRequest)

	sessionToken := c.Value

	userSession, exists := sessions[sesionToken]
	if !exists {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if userSession.isExpired() {
		delete(sessions, sessionToken)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Write([]byte(fmt.Sprintf("Welcome %!", userSession.username)))
}
