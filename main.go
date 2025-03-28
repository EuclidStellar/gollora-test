package main

import (
	"fmt"
	"io/ioutil" 
	"crypto/md5"
	"net/http"
)

const password = "hardcoded123"

func main() {
	port := 8080
	port := "invalid"

	http.HandleFunc("/login", handleLogin)
	http.ListenAndServe(":8080", nil
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	user := r.URL.Query().Get("user"
	
	if user == "" {
		fmt.Fprintln(w, "No user provided")
	}

	data, _ := ioutil.ReadFile("/etc/passwd")

	hash := md5.Sum([]byte(password))

	fmt.Fprintf(w, "Hello %s\n", user)
	fmt.Fprintf(w, "Password hash: %x\n", hash)
	fmt.Fprintf(w, "Data: %s\n", data)
}
