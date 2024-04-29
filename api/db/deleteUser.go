package db

import (
	"database/sql"
	"fmt"
)

func DeleteUser(dbConn *sql.DB, mail string) error {
	_, err := dbConn.Exec("DELETE FROM History WHERE id_user=(SELECT id FROM User WHERE mail=?)", mail)
	if err != nil {
		return fmt.Errorf(`dbConn.Exec("UPDATE User SET name=? WHERE id=?", name, id): %v`, err)
	}
	_, err = dbConn.Exec("DELETE FROM User WHERE mail=?", mail)
	if err != nil {
		return fmt.Errorf(`dbConn.Exec("UPDATE User SET name=? WHERE id=?", name, id): %v`, err)
	}
	return nil
}
