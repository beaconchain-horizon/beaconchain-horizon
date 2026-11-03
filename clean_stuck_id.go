package main

import (
	"database/sql"
	"fmt"
	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "horizon-dev.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// حذف رکورد گیر کرده با ID 1555
	res, err := db.Exec("DELETE FROM license_records WHERE id = 1555")
	if err != nil {
		fmt.Println("Error deleting:", err)
	} else {
		rows, _ := res.RowsAffected()
		fmt.Printf("✅ Deleted stuck record (rows affected: %d)\n", rows)
	}
}
