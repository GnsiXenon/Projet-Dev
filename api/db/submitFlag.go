package db

import (
	"database/sql"
	"fmt"
)

func SubmitFlag(dbConn *sql.DB, userId int, challId int, flag string) (bool, error) {
	result := false
	resultRow := 0
	row := dbConn.QueryRow("SELECT COUNT(*) FROM Challenge WHERE id=? AND flag=?", challId, flag)
	if err := row.Scan(&resultRow); err != nil {
		return result, fmt.Errorf("row.Scan(&resultRow): %v", err)
	}
	if resultRow > 0 {
		result = true
		score := 0
		row2 := dbConn.QueryRow("SELECT score FROM User WHERE id=?", userId)
		if err := row2.Scan(&score); err != nil {
			return result, fmt.Errorf("row.Scan(&resultRow): %v", err)
		}
		if _, err := dbConn.Exec("UPDATE User SET score=? WHERE id=?", score+100, userId); err != nil {
			return result, fmt.Errorf(`dbConn.Exec("UPDATE User SET score=? WHERE id=?", score+100, userId): %v`, err)
		}
	}
	return result, nil
}
