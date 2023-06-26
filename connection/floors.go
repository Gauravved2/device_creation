package connection

import (
	"database/sql"
	"time"
)

func AddFloors(db *sql.DB, code string, bcode string, name string , shortName string) error {
	creationTime := time.Now()
	stmt, err := db.Prepare("INSERT INTO floors(created_at, updated_at, code, building_code, name, short_name, physical_index) VALUES(?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	res, err := stmt.Exec(creationTime, creationTime, code, bcode, name, shortName, 0)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}