package connection

import (
	"database/sql"
	"time"
)

func AddBuilding(db *sql.DB, buildingCode string, name string, shortName string) (error){
	creationTime := time.Now()
	stmt, err := db.Prepare("INSERT INTO buildings(created_at, updated_at, code, name, short_name) VALUES(?,?,?,?,?)")
	if err!=nil {
		return err
	}
	res, err := stmt.Exec(creationTime, creationTime, buildingCode, name, shortName)
	if err!=nil{
		return err
	}
	_, err = res.RowsAffected()
	if err!=nil{
		return err
	}
	return nil
}