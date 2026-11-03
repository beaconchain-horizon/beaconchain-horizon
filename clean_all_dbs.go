package main

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	_ "modernc.org/sqlite"
)

func main() {
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if filepath.Ext(path) == ".db" {
			db, err := sql.Open("sqlite", path)
			if err == nil {
				// حذف رکورد 1555
				_, delErr := db.Exec("DELETE FROM license_records WHERE id = 1555")
				if delErr == nil {
					fmt.Printf("✅ Cleaned ID 1555 from: %s\n", path)
				}
				// حذف لایسنس قبلی اگر وجود داشته باشد تا دوباره ثبت شود
				db.Exec("DELETE FROM license_records WHERE license_id = 'HORIZON-GOLD-005'")
				db.Close()
			}
		}
		return nil
	})
}
