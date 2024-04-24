package db

import (
	"database/sql"
	"fmt"
)

func GetUsers(dbConn *sql.DB) ([]User, error) {
	rows, err := dbConn.Query("SELECT id, name, score FROM User")
	if err != nil {
		return nil, fmt.Errorf(`dbConn.Query("SELECT (id, name, score) FROM User"): %v`, err)
	}
	result := []User{}
	for rows.Next() {
		user := User{}
		rows.Scan(&user.Id, &user.Name, &user.Score)
		result = append(result, user)
	}
	return result, nil
}
