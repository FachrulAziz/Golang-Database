package belajargolangdatabase

import (
	"database/sql"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/belajar_golang_database?parseTime=true")
	if err != nil {
		panic(err)
	}
	// db.SetMaxIdleConns(10)                  // Pengaturan berapa jumlah koneksi minimal yang dibuat
	// db.SetMaxOpenConns(100)                 // Pengaturan berapa jumlah koneksi maksimal yang dibuat
	// db.SetConnMaxIdleTime(5 * time.Minute)  // Pengaturan berapa lama koneksi yang sudah tidak digunakan akan dihapus
	// db.SetConnMaxLifetime(60 * time.Minute) // Pengaturan berapa lama koneksi boleh digunakan

	return db
}
