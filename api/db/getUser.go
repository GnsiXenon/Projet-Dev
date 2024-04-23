package db

import (
	"database/sql"
	"fmt"
)

func GetUser(dbConn *sql.DB, mail string) (*User, error) {
	row := dbConn.QueryRow("SELECT (id, name, mail, score) FROM User WHERE mail=?", mail)
	result := User{}
	if err := row.Scan(&result); err != nil {
		return nil, fmt.Errorf("row.Scan(&result): %v", err)
	}
	return &result, nil
}
