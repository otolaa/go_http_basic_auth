package main

import (
	"fmt"
	"log"
	"net/http"
)

var users = map[string]string{
	"username": "password",
}

func isAuthorised(username string, password string) bool {
	pass, ok := users[username]
	if !ok {
		return false
	}

	return password == pass
}

func httpBasicAuth(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	username, password, ok := r.BasicAuth()
	if !ok {
		w.Header().Add("WWW-Authenticate", `Basic realm="Give username and password"`)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"message": "No basic auth present 🌶️"}`))
		return
	}

	if !isAuthorised(username, password) {
		w.Header().Add("WWW-Authenticate", `Basic realm="Give username and password"`)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"message": "Invalid username or password 🌵"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"meesage": "welcome 🌴"}`))
}

func main() {
	http.HandleFunc("/", httpBasicAuth)

	fmt.Println("Starting Server at port :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
