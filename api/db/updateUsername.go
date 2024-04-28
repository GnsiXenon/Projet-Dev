package db

import (
	"database/sql"
	"fmt"
)

// UpdateUsername update the username of a user in the database
func UpdateUsername(dbConn *sql.DB, id *int, name string) error {
	_, err := dbConn.Exec("UPDATE User SET name=? WHERE id=?", name, id)
	if err != nil {
		return fmt.Errorf(`dbConn.Exec("UPDATE User SET name=? WHERE id=?", name, id): %v`, err)
	}
	return nil
}
