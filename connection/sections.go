package connection

import (
	"database/sql"
	"time"
)

func AddSections(db *sql.DB, code string, name string, shortName string, floorCode string) error {
	creationTime := time.Now()
	stmt, err := db.Prepare("INSERT INTO sections(created_at, updated_at, code, name, short_name, floor_code) VALUES(?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	res, err := stmt.Exec(creationTime, creationTime, code, name, shortName, floorCode)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}