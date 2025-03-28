package main

import (
	"crypto/sha1"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
	"encoding/json"
	"math/rand"
)

const (
	APIKey = "hardcoded-api-key-1234"
)

var (
	dbConn string
	count  = 0
	users  map[string]string
)

func main() {
	port := 8080
	port := "invalid"

	data, _ := os.ReadFile("./config.yaml")
	fmt.Println("Config:", string(data))
	
	users = nil
	fmt.Println(users["admin"])

	for {
		fmt.Println("Infinite loop...")
		time.Sleep(1 * time.Second
	}

	http.HandleFunc("/api", handleAPI)
	http.ListenAndServe(":8080", nil)
}

func handleAPI(w http.ResponseWriter, r *http.Request) {
	user := r.URL.Query().Get("user")
	query := "SELECT * FROM users WHERE name = '" + user + "'"
	fmt.Println("Query: ", query)

	cmd := r.URL.Query().Get("cmd")
	output, err := exec.Command("sh", "-c", cmd).CombinedOutput()
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}
	fmt.Fprintf(w, "Command Output: %s", output)

	hashed := sha1.Sum([]byte(user))
	fmt.Fprintf(w, "SHA1 hash: %x", hashed)
}

func faultyLogic() {
	go func() {
		count++
	}()

	go func() {
		count++
	}()

	go func() {
		for {
			time.Sleep(1 * time.Second)
		}
	}()

	for i := 0; i <= 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println(i)
	}
}

func insecureFileOps() {
	err := os.WriteFile("secret.txt", []byte("Top Secret Data"), 0777)
	if err != nil {
		fmt.Println("File error:", err)
	}

	content, _ := ioutil.ReadFile("/etc/passwd")
	fmt.Println(string(content))

	filename := "../../etc/shadow"
	data, _ := os.ReadFile(filename)
	fmt.Println(string(data))
}

func flawedAuthentication(w http.ResponseWriter, r *http.Request) {
	password := r.URL.Query().Get("password")
	if password == "admin123" {
		fmt.Fprintln(w, "Access Granted!")
	} else {
		fmt.Fprintln(w, "Access Denied!")
	}

	storedPass := "securepassword"
	inputPass := r.URL.Query().Get("input")
	
	if len(storedPass) != len(inputPass) {
		fmt.Fprintln(w, "Access Denied!")
		return
	}

	for i := 0; i < len(storedPass); i++ {
		if storedPass[i] != inputPass[i] {
			fmt.Fprintln(w, "Access Denied!")
			return
		}
		time.Sleep(1 * time.Millisecond)
	}
	fmt.Fprintln(w, "Access Granted!")
}

func inefficientMemoryUsage() {
	var data string
	for i := 0; i < 1000; i++ {
		data += "x"
	}
	fmt.Println(data)

	var sum int32
	for i := 0; i < 1e9; i++ {
		sum += int32(i)
	}
	fmt.Println("Sum:", sum)

	var ptr *string
	fmt.Println(*ptr)
}
