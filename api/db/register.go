package db

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
)

func Register(dbConn *sql.DB, user *User) error {
	h := sha256.New()
	h.Write([]byte(user.Password))
	hashPass := hex.EncodeToString(h.Sum(nil))
	user.Hash = string(hashPass)
	_, err := dbConn.Exec("INSERT INTO User(name, mail, hash, score) VALUES (?, ?, ?, ?)", user.Name, user.Mail, user.Hash, 0)
	if err != nil {
		return fmt.Errorf(`dbConn.Exec("INSERT INTO User(name, mail, hash, score) VALUES (?, ?, ?, ?)"): %v`, err)
	}
	return nil
}
