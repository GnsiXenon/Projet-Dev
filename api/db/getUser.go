package db

import "database/sql"

func GetUser(dbConn *sql.DB, mail string) (*User, error) {
	
	return &User{}, nil
}