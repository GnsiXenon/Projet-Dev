package db

import (
	"database/sql"
	"fmt"
)

func SubmitFlag(dbConn *sql.DB, userId int, challId int, flag string) (int, error) {
	result := -1
	histoRes := 0
	histoRow := dbConn.QueryRow("SELECT COUNT(*) FROM History WHERE id_user=? AND id_challenge=?", userId, challId)
	if err := histoRow.Scan(&histoRes); err != nil {
		return result, fmt.Errorf("histoRow.Scan(&histoRes): %v", err)
	}
	if histoRes != 0 {
		return 0, nil
	}
	resultRow := 0
	row := dbConn.QueryRow("SELECT COUNT(*) FROM Challenge WHERE id=? AND flag=?", challId, flag)
	if err := row.Scan(&resultRow); err != nil {
		return result, fmt.Errorf("row.Scan(&resultRow): %v", err)
	}
	if resultRow > 0 {
		result = 1
		score := 0
		row2 := dbConn.QueryRow("SELECT score FROM User WHERE id=?", userId)
		if err := row2.Scan(&score); err != nil {
			return result, fmt.Errorf("row.Scan(&resultRow): %v", err)
		}
		if _, err := dbConn.Exec("UPDATE User SET score=? WHERE id=?", score+100, userId); err != nil {
			return result, fmt.Errorf(`dbConn.Exec("UPDATE User SET score=? WHERE id=?", score+100, userId): %v`, err)
		}
		if _, err := dbConn.Exec("INSERT INTO History (id_user, id_challenge) VALUES (?, ?)", userId, challId); err != nil {
			return -1, fmt.Errorf(`dbConn.Exec("INSERT INTO History (id_user, id_challenge) VALUES (?, ?)", userId, challId): %v`, err)
		}
	}
	return result, nil
}
