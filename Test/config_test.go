package Test

import (
	"fmt"
	"github.com/spf13/viper"
	"testing"
)

func testconfig(t *testing.T) {
	// 从配置文件中获取数据库配置
	viper.SetConfigFile("../config.yaml") // 配置文件的名称和路径
	dbConfig := viper.Sub("database")
	username := dbConfig.GetString("username")
	password := dbConfig.GetString("password")
	host := dbConfig.GetString("host")
	port := dbConfig.GetInt("port")
	dbname := dbConfig.GetString("dbname")

	// 打印配置信息
	fmt.Println("Database Configuration:")
	fmt.Println("Username:", username)
	fmt.Println("Password:", password)
	fmt.Println("Host:", host)
	fmt.Println("Port:", port)
	fmt.Println("DB Name:", dbname)
}
