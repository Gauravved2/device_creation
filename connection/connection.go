package connection

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func CreateConnection() (*sql.DB, error) {
	dbUser := "root"
	dbPass := "thisismysql"
	dbPort := "3306"
	dbHost := "localhost"
	dbName := "ResourceManagement"
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	dbConnection, err := sql.Open("mysql", connString)
	if err != nil {
		return nil, err
	}

	// err = dbConnection.Ping()
	// if err!=nil {
	// 	return nil, err
	// }
	fmt.Println("Returning data")
	return dbConnection, nil

}
