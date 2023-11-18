package main

import (
	"fmt"
	"log"
	"net/http"
	"yiliao/Database"
	"yiliao/Handler"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

// 定义JWT密钥

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Content-Disposition, Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Disposition,Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
func main() {
	//启用gin 框架
	r := gin.Default()
	r.Use(Cors())
	r.Use(Handler.RouterAuth)
	//加载配置文件
	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
	//连接数据库
	dbConfig := viper.Sub("database")
	db, err := Database.ConnectToDatabase(dbConfig.GetString("username"), dbConfig.GetString("password"), dbConfig.GetString("host"), dbConfig.GetInt("port"), dbConfig.GetString("dbname"))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	//开启服务
	//登录
	r.POST("/login", Handler.LoginHandler(db))
	//获取所有用户
	r.GET("/user", Handler.GetUsersHandler(db))
	// 创建策略
	r.POST("/stragety/create", Handler.CreateStragety(db))
	//获取用户所拥有的文件的列表
	r.GET("/file", Handler.GetFileListHandler(db))
	//下载文件
	r.GET("/file/download", Handler.DownloadFileHandler(db))
	r.POST("/file/share", Handler.ShareFileHandler(db))
	r.GET("/file/share", Handler.GetShareHandler(db))
	r.Run(":8080")
}
