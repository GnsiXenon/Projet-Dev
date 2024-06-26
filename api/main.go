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
			defer dbConn.Close()
			isLog, err := db.Login(dbConn, user.Mail, user.Password)
			if err != nil {
				log.Printf("db.Login(dbConn, user.Name, user.Password): %v", err)
				http.Error(w, fmt.Sprintf("db.Login(dbConn, user.Name, user.Password): %v", err), http.StatusInternalServerError)
				return
			}
			if isLog {
				theUser, err := db.GetUser(dbConn, user.Mail)
				if err != nil {
					log.Printf("db.GetUser(dbConn, user.Mail): %v", err)
					http.Error(w, fmt.Sprintf("db.GetUser(dbConn, user.Mail): %v", err), http.StatusInternalServerError)
					return
				}
				byteUser, err := json.MarshalIndent(theUser, "", "	")
				if err != nil {
					log.Printf(`json.MarshalIndent(theUser, "", "	"): %v`, err)
					http.Error(w, fmt.Sprintf(`json.MarshalIndent(theUser, "", "	"): %v`, err), http.StatusInternalServerError)
					return
				}
				w.Write(byteUser)
			} else {
				w.Write([]byte(`{"error":"wrong credentials"}`))
			}
		} else {
			http.Error(w, fmt.Sprintf("Wants request method POST, got : %s\n", r.Method), http.StatusBadRequest)
			return
		}
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
			defer dbConn.Close()
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
		if r.Method == http.MethodPost {
			var data struct {
				UserId   int    `json:"user-id"`
				ChallId  int    `json:"chall-id"`
				Flag     string `json:"flag"`
				UserMail string `json:"mail"`
			}
			if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
				log.Printf("json.NewDecoder(r.Body).Decode(&data): %v", err)
				http.Error(w, fmt.Sprintf("json.NewDecoder(r.Body).Decode(&data): %v", err), http.StatusInternalServerError)
				return
			}
			dbConn, err := db.GetConn()
			if err != nil {
				log.Printf("db.GetConn(): %v", err)
				http.Error(w, fmt.Sprintf("db.GetConn(): %v", err), http.StatusInternalServerError)
				return
			}
			defer dbConn.Close()
			_, err = db.SubmitFlag(dbConn, data.UserId, data.ChallId, data.Flag)
			if err != nil {
				log.Printf("db.SubmitFlag(dbConn, data.UserId, data.ChallId, data.Flag): %v", err)
				http.Error(w, fmt.Sprintf("db.SubmitFlag(dbConn, data.UserId, data.ChallId, data.Flag): %v", err), http.StatusInternalServerError)
				return
			}
			theUser, err := db.GetUser(dbConn, data.UserMail)
			if err != nil {
				log.Printf("db.GetUser(dbConn, user.Mail): %v", err)
				http.Error(w, fmt.Sprintf("db.GetUser(dbConn, user.Mail): %v", err), http.StatusInternalServerError)
				return
			}
			byteUser, err := json.MarshalIndent(theUser, "", "	")
			if err != nil {
				log.Printf(`json.MarshalIndent(theUser, "", "	"): %v`, err)
				http.Error(w, fmt.Sprintf(`json.MarshalIndent(theUser, "", "	"): %v`, err), http.StatusInternalServerError)
				return
			}
			w.Write(byteUser)
		} else {
			http.Error(w, fmt.Sprintf("Wants request method POST, got : %s\n", r.Method), http.StatusBadRequest)
			return
		}
	})
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			dbConn, err := db.GetConn()
			if err != nil {
				log.Printf("db.GetConn(): %v", err)
				http.Error(w, fmt.Sprintf("db.GetConn(): %v", err), http.StatusInternalServerError)
				return
			}
			defer dbConn.Close()
			users, err := db.GetUsers(dbConn)
			if err != nil {
				log.Printf("db.GetUsers(dbConn): %v", err)
				http.Error(w, fmt.Sprintf("db.GetUsers(dbConn): %v", err), http.StatusInternalServerError)
				return
			}
			byteUsers, err := json.MarshalIndent(users, "", "	")
			if err != nil {
				log.Printf(`json.MarshalIndent(users, "", "	"): %v`, err)
				http.Error(w, fmt.Sprintf(`json.MarshalIndent(users, "", "	"): %v`, err), http.StatusInternalServerError)
				return
			}
			w.Write(byteUsers)
		} else {
			http.Error(w, fmt.Sprintf("Wants request method GET, got : %s\n", r.Method), http.StatusBadRequest)
			return
		}
	})
	http.HandleFunc("/delete-user", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			var data struct {
				Mail string `jso:"mail"`
			}
			if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
				log.Printf("json.NewDecoder(r.Body).Decode(&data): %v", err)
				http.Error(w, fmt.Sprintf("json.NewDecoder(r.Body).Decode(&data): %v", err), http.StatusInternalServerError)
				return
			}
			dbConn, err := db.GetConn()
			if err != nil {
				log.Printf("db.GetConn(): %v", err)
				http.Error(w, fmt.Sprintf("db.GetConn(): %v", err), http.StatusInternalServerError)
				return
			}
			defer dbConn.Close()
			if err := db.DeleteUser(dbConn, data.Mail); err != nil {
				log.Printf("db.DeleteUser(dbConn, data.Mail): %v", err)
				http.Error(w, fmt.Sprintf("db.DeleteUser(dbConn, data.Mail): %v", err), http.StatusInternalServerError)
				return
			}
		} else {
			http.Error(w, fmt.Sprintf("Wants request method POST, got : %s\n", r.Method), http.StatusBadRequest)
			return
		}
	})
	// Start the API
	log.Printf("API is up on port 0.0.0.0:%s 🔥 : http://localhost:%s", ports, ports)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", ports), nil); err != nil {
		log.Printf(`http.ListenAndServe(fmt.Sprintf(":%%s", ports), nil): %v`, err)
		return
	}
}
