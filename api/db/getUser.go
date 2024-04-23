package db

import (
	"database/sql"
	"fmt"
)

func GetUser(dbConn *sql.DB, mail string) (*User, error) {
	row, err := dbConn.Query("SELECT id, name, mail, score FROM User WHERE mail=?", mail)
	if err != nil {
		return nil, fmt.Errorf(`dbConn.Query("SELECT (id, name, mail, score) FROM User WHERE mail=?", mail): %v`, err)
	}
	result := User{}
	for row.Next() {
		if err := row.Scan(&result.Id, &result.Name, &result.Mail, &result.Score); err != nil {
			return nil, fmt.Errorf("row.Scan(&result.Id, &result.Name, &result.Mail, &result.Score): %v", err)
		}
	}
	return &result, nil
}
