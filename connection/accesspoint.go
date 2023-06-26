package connection

import (
	"database/sql"
	"time"
)

func AddAccessPoint(db *sql.DB, code string, name string, BSSID string, SSID string, floorCode string, sectionCode string) error {
	creationTime := time.Now()
	stmt, err := db.Prepare("INSERT INTO access_points(created_at, updated_at, code, name, bss_id, ss_id, floor_code, section_code, is_perimeter) VALUES(?,?,?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	res, err := stmt.Exec(creationTime, creationTime, code, name, BSSID, SSID, floorCode, sectionCode, 0)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}