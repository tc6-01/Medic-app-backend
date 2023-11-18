package yiliao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"log"
	"testing"
	"yiliao/Database"
)

var sault string = "YiLiao"

func TestName(t *testing.T) {
	var db *sql.DB
	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
	dbConfig := viper.Sub("database")
	db, _ = Database.ConnectToDatabase(dbConfig.GetString("username"), dbConfig.GetString("password"), dbConfig.GetString("host"), dbConfig.GetInt("port"), dbConfig.GetString("dbname"))

	defer db.Close()
	var storedPassword string
	/*使用盐值加密密码*/
	conn_err := db.QueryRow("SELECT password FROM users WHERE user_name = ?", "admin").Scan(&storedPassword)
	if conn_err != nil {
		print(conn_err)
	}
	print(storedPassword)
}
