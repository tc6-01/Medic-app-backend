package Handler

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ShareFile struct {
	Expire    int64  `json:"expire"`
	FileName  string `json:"fileName"`
	Target    string `json:"target"`
	Use_count int64  `json:"use"`
	UseLimit  int64  `json:"useLimit"`
	IsGroup   int64  `json:"isGroup"'`
	FileSize  int64  `json:"fileSize"`
}

// 获取当前用户共享的所有文件
func GetShareHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从上下文中获取用户信息
		user, exists := c.Get("user")
		if !exists {
			handleError(c, "Parser Error", "User not found in context")
			return
		}

		// 执行 SQL 查询以获取当前用户共享的所有文件
		rows, err := db.Query(`
            SELECT file_name, target, expire, use_limit,is_group,use_count,file_size
            FROM file_share_list
            WHERE owner = ?`, user)

		if err != nil {
			handleError(c, "DB Error", "Database query error")
			return
		}
		defer rows.Close()

		// 创建一个切片来存储共享文件的信息
		var sharedFiles []ShareFile

		// 遍历查询结果并填充 sharedFiles 切片
		for rows.Next() {
			var (
				expire    int64
				fileName  string
				target    string
				use_count int64
				useLimit  int64
				isGroup   int64
				fileSize  int64
			)

			err := rows.Scan(&fileName, &target, &expire, &useLimit, &isGroup, &use_count, &fileSize)
			if err != nil {
				handleError(c, "DB Error", "Error scanning rows")
				return
			}

			// 创建共享文件对象并填充数据
			sharedFile := ShareFile{
				Expire:    expire,
				FileName:  fileName,
				Target:    target,
				Use_count: use_count,
				UseLimit:  useLimit,
				IsGroup:   isGroup,
				FileSize:  fileSize,
			}

			sharedFiles = append(sharedFiles, sharedFile)
		}

		// 检查是否有错误
		if err := rows.Err(); err != nil {
			handleError(c, "DB Error", "Error in rows")
			return
		}

		// 返回共享文件列表
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "Success", "data": sharedFiles})
	}
}
