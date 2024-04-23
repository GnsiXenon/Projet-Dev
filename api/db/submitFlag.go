package db

import "database/sql"

func SubmitFlag(dbConn *sql.DB, userId int, challId int, flag string) (bool, error) {
	result := false
	return result, nil
}
