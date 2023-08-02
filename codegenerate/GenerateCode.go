package codegenerate

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)


func Connection() *sql.DB {
	dbUser := "root"
	dbPass := "thisismysql"
	dbPort := "3306"
	dbHost := "localhost"
	dbName := "ResourceManagement"
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	dbConnection, err := sql.Open("mysql", connString)
	if err != nil {
		panic(err.Error())
	}
	return dbConnection
}
func GenerateCode(dbCon *sql.DB, table string, prefix string) string {
	c := getLatestId(dbCon,table, prefix)
	code := fmt.Sprintf("%s%d",prefix,c)
	fmt.Println(code)
	// stmt, err := dbCon.Prepare("INSERT INTO "+table+" VALUES(?,?)")
	// if err!=nil{
	// 	panic(err)
	// }
	// _, err = stmt.Exec(code, name)
	// if err!=nil{
	// 	panic(err.Error())
	// }
	// fmt.Println("Inserted Successfully")
	return code
}
func getLatestId(dbCon *sql.DB, table string, prefix string) int {
	query := "SELECT id FROM "+table+" ORDER BY id desc LIMIT 1"
	var col1 string 
	var code int = 0
	err := dbCon.QueryRow(query).Scan(&col1)
	if err!=nil{
		if(err.Error() == "sql: no rows in result set"){
			code = 1
		} else{
			panic(err.Error())
		}
	} else{
		stringcode := strings.TrimPrefix(col1, prefix)
		code, err = strconv.Atoi(stringcode)
		if err!=nil{
			panic(err.Error())
		}
		code ++
	}
	return code
}
