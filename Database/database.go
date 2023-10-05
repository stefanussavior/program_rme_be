package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

var DB *sql.DB
var err error

func ConnectDatabase() {

	viper.SetConfigFile("env")
	viper.SetConfigName("development")
	viper.AddConfigPath("./env")

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	DBUsername := viper.GetString("MYSQL_USERNAME")
	DBPassword := viper.GetString("MYSQL_PASSWORD")
	DBHost := viper.GetString("MYSQL_HOST")
	DBPort := viper.GetString("MYSQL_PORT")
	DBName := viper.GetString("MYSQL_DBNAME")

	ConnectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DBUsername, DBPassword, DBHost, DBPort, DBName)

	DB, err = sql.Open("mysql", ConnectionString)

	if err != nil {
		panic(err.Error())
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	fmt.Println("Database berhasil dijalankan")
}
