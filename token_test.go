package yiliao

import (
	"database/sql"
	"log"
	"testing"
	"yiliao/Database"
	"yiliao/Handler"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func RegisterTest(db *sql.DB) {
	/*使用盐值加密密码*/
	password := "123456"
	err := Handler.RegisterUser(db, "admin", password)
	if err != nil {
		log.Println(err.Error())
	}
}
func LoginTest(db *sql.DB) {
	conn_err := db.QueryRow("SELECT password FROM user WHERE user_name = ?", "admin")
	if conn_err != nil {
		print(conn_err)
	}
}
func Test(t *testing.T) {
	var db *sql.DB
	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
	dbConfig := viper.Sub("database")
	db, _ = Database.ConnectToDatabase(dbConfig.GetString("username"), dbConfig.GetString("password"), dbConfig.GetString("host"), dbConfig.GetInt("port"), dbConfig.GetString("dbname"))

	defer db.Close()
	RegisterTest(db)
	LoginTest(db)
}
