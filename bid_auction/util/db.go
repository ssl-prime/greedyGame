package util

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

//ConnectDB ...
func ConnectDB() (*sql.DB, error) {
	var (
		dbUser     = "root"
		dbPassword = "spatico"
		dbHost     = "localhost"
		dbPort     = "3306"
		dbName     = "greddy"
	)
	db, err := sql.Open("mysql", dbUser+":"+dbPassword+"@tcp("+dbHost+":"+dbPort+")/"+dbName)
	return db, err
}
