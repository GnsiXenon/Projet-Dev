package db

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Return the connection with the database
func GetConn() (*sql.DB, error) {
	dbHostname := "localhost"
	envDbHostname, isSet := os.LookupEnv("DB_HOSTNAME")
	if isSet {
		dbHostname = envDbHostname
	}
	db, err := sql.Open("mysql", fmt.Sprintf("database-user:1e40ebe3f2e384cd625a50399857efdc32d4b7031bea706634816adbc775722b@tcp(%s:3306)/db", dbHostname))
	if err != nil {
		return nil, fmt.Errorf(`sql.Open("mysql", fmt.Sprintf("database-user:1e40ebe3f2e384cd625a50399857efdc32d4b7031bea706634816adbc775722b@tcp(%%s:3306)/db", dbHostname)): %v`, err)
	}
	// test the connection with the database
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("db.Ping(): %v", err)
	}
	// Setup the database connection
	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(5)
	return db, nil
}
