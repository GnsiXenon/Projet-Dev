package db

import (
	"database/sql"
	"fmt"
)

// DeleteUser delete a user from the database
func DeleteUser(dbConn *sql.DB, id *int) error {
	_, err := dbConn.Exec("DELETE FROM User WHERE id-user=?", id)
	if err != nil {
		return fmt.Errorf(`dbConn.Exec("DELETE FROM User WHERE id-user=?", id): %v`, err)
	}
	return nil

}
