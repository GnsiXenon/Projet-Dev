package db

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
)

func Register(dbConn *sql.DB, user *User) error {
	// Check if user Already Exists
	row := dbConn.QueryRow("SELECT COUNT(*) FROM User WHERE mail=?", user.Mail)
	result := 0
	if err := row.Scan(&result); err != nil {
		return fmt.Errorf("row.Scan(&result): %v", err)
	}
	if result != 0 {
		return fmt.Errorf("User already exist")
	}
	// Calculate Password hash
	h := sha256.New()
	h.Write([]byte(user.Password))
	hashPass := hex.EncodeToString(h.Sum(nil))
	user.Hash = string(hashPass)
	// Create user
	_, err := dbConn.Exec("INSERT INTO User(name, mail, hash, score) VALUES (?, ?, ?, ?)", user.Name, user.Mail, user.Hash, 0)
	if err != nil {
		return fmt.Errorf(`dbConn.Exec("INSERT INTO User(name, mail, hash, score) VALUES (?, ?, ?, ?)"): %v`, err)
	}
	return nil
}
