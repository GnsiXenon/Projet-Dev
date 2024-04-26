package db

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
)

func Login(dbConn *sql.DB, mail string, password string) (bool, error) {
	h := sha256.New()
	h.Write([]byte(password))
	hashPass := hex.EncodeToString(h.Sum(nil))
	row := dbConn.QueryRow("SELECT COUNT(*) FROM User WHERE mail=? AND hash=?", mail, hashPass)
	result := 0
	if err := row.Scan(&result); err != nil {
		return false, fmt.Errorf("row.Scan(&result): %v", err)
	}
	return result != 0, nil
}
