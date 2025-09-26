package main

import (
	"crypto/subtle"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	addr          = ":8080"
	hardUser      = "testuser"
	hardPassword  = "dummyPassword123!" // <-- hardcoded dummy password for testing
	sessionCookie = "session_token"
)

// simple JSON response helper
type jsonResp struct {
	OK      bool   `json:"ok"`
	Message string `json:"message,omitempty"`
}

// constantTimeEquals compares two strings in constant time to avoid timing attacks.
func constantTimeEquals(a, b string) bool {
	if len(a) != len(b) {
		// still run a constant-time compare on equal-length slices to avoid leaking length info
		// create a dummy slice of same length as a
		_ = subtle.ConstantTimeCompare([]byte(a), []byte(a))
		return false
	}
	return subtle.ConstantTimeCompare([]byte(a), []byte(b)) == 1
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/login", loginHandler)       // handles form POST
	http.HandleFunc("/api/login", apiLoginHandler) // handles JSON login (POST)
	http.HandleFunc("/protected", protectedHandler)
	http.HandleFunc("/logout", logoutHandler)

	log.Printf("Starting server at http://localhost%s\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}

// indexHandler serves a tiny HTML login form.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	html := `<!doctype html>
<html>
<head><meta charset="utf-8"><title>Test Login</title></head>
<body>
<h2>Test Login (hardcoded creds)</h2>
<form action="/login" method="POST">
  <label>Username: <input name="username" /></label><br/><br/>
  <label>Password: <input type="password" name="password" /></label><br/><br/>
  <button type="submit">Login</button>
</form>
<p>Or use the JSON endpoint POST /api/login with body {"username":"testuser","password":"dummyPassword123!"}</p>
</body>
</html>`
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, html)
}

// loginHandler processes the HTML form login.
func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	if authenticate(username, password) {
		// set a simple cookie for demo purposes
		http.SetCookie(w, &http.Cookie{
			Name:     sessionCookie,
			Value:    "authenticated", // in real apps set a secure session id
			Path:     "/",
			Expires:  time.Now().Add(30 * time.Minute),
			HttpOnly: true,
			Secure:   false, // set to true when using HTTPS
		})
		http.Redirect(w, r, "/protected", http.StatusSeeOther)
		return
	}

	// fail
	http.Error(w, "invalid credentials", http.StatusUnauthorized)
}
