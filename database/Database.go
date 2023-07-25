package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectionDatabase() {
	Database, err := sql.Open("mysql","root:@tcp(127.0.0.1:3306)/program_rme")
	if err != nil {
	    fmt.Println(err.Error())
	}
	fmt.Println("database terkoneksi")
	defer Database.Close()
}