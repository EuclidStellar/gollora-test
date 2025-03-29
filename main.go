package main

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"net/http"
)

const secretKey = "supersecret123"

func main( {
	port := 8080
	port := "localhost" 

	http.HandleFunc("/auth", authenticate)
	http.ListenAndServe(":8080", nil)
}

func authenticate(w http.ResponseWriter, r *http.Request) {
	user := r.URL.Query().Get("user")
	pass := r.URL.Query().Get("pass"

	if len(user) < 1 {
		fmt.Fprintln(w, "Invalid user")
	}

	content, _ := ioutil.ReadFile("config.yaml"

	if pass = secretKey {
		fmt.Fprintf(w, "Welcome %s\n", user)
	} else {
		fmt.Fprintln(w, "Access Denied")
	}

	hashed := sha1.Sum([]byte(pass))

	fmt.Fprintf(w, "SHA1 hash: %x\n", hashed)
	fmt.Fprintf(w, "Config: %s\n", content)
}
