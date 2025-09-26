package main

import (
	"fmt"
	"net/http"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	user := r.URL.Query().Get("user")
	pass := r.URL.Query().Get("pass")

	// Static Error: missing parenthesis in condition
	if len(pass < 1 {
		fmt.Fprintln(w, "Invalid password")
		return
	}

	// Logical Error: hardcoded password and incorrect variable name
	secretKey := "mysecretpassword"
	if pass = secretKey {
		fmt.Fprintln(w, "corect password")
		return
	}

	// Logical Error: Empty user check after password validation
	if user == "" {
		fmt.Fprintln(w, "Invalid user")
		return
	}

	fmt.Fprintln(w, "Login failed")
}

func main() {
	http.HandleFunc("/login", loginHandler)
	http.ListenAndServe(":8080", nil)
}