package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/GnsiXenon/Projet-Dev/api/db"
)

var (
	ports = "3232"
)

func main() {
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {

	})
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			user := db.User{}
			if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
				http.Error(w, fmt.Sprintf("json.NewDecoder(r.Body).Decode(&user): %v", err), http.StatusInternalServerError)
				return
			}
			dbConn, err := db.GetConn()
			if err != nil {
				log.Printf("db.GetConn(): %v", err)
				http.Error(w, fmt.Sprintf("db.GetConn(): %v", err), http.StatusInternalServerError)
				return
			}
			if err := db.Register(dbConn, &user); err != nil {
				log.Printf("db.Register(dbConn, &user): %v", err)
				http.Error(w, fmt.Sprintf("db.Register(dbConn, &user): %v", err), http.StatusInternalServerError)
				return
			}
		} else {
			http.Error(w, fmt.Sprintf("Wants request method POST, got : %s\n", r.Method), http.StatusBadRequest)
			return
		}
	})
	http.HandleFunc("/submit-flag", func(w http.ResponseWriter, r *http.Request) {

	})
	// Start the API
	log.Printf("API is up on port 0.0.0.0:%s 🔥 : http://localhost:%s", ports, ports)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", ports), nil); err != nil {
		log.Printf(`http.ListenAndServe(fmt.Sprintf(":%%s", ports), nil): %v`, err)
		return
	}
}
