package db

import (
	"database/sql"
)

func Login(dbConn *sql.DB, username string, password string) (bool, error) {
	return true, nil
}
