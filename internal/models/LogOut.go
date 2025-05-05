package main

import (
	"net/http"
	"time"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnautthorized)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		return

	}

	sessionToken := c.Value

	delete(sessions, sessionToken)

	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   "",
		Expires: time.Now(),
	})
}
