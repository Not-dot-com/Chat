package main

import (
	"encoding/json"
	"net/http"
)

func Signin(w http.ResponseWrter, r *http.Requset) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nill {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expectedPassword, ok := users[creds.Username]

	if !ok || expectedPassword != creds.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	sessionToken := uuid.NewSting()
	expiresAt := time.Now().(120 * time.Second)

	sessions[sessionToken] = session {
		username: creds.Username,
		expiry: expiresAt,
	}

	http.SetCookie(w, &http.Cookie{
		Name: "session_token",
		Value: sessionToken,
		Expires: expiresAt,
	})
}
