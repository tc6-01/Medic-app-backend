package Handler

import (
	"database/sql"
	"net/http"
	"yiliao/Utils"

	"github.com/gin-gonic/gin"
)

/*CREATE TABLE IF NOT EXISTS file_share_list (
	expire BIGINT,
	file_name VARCHAR(255),
	file_size BIGINT,
	owner VARCHAR(255),
	use_count BIGINT DEFAULT 0,
	use_limit BIGINT,
	target VARCHAR(255),
	is_group TINYINT ,
);*/
// 共享文件 handler
// 客户端传入参数为fileName,target,expire,useLimit,isGroup
// 服务器段根据fileName 从file_list中查到 file_size, owner
// 服务器 从c中获得user ，若user等于target则报错(不能分享给自己)
// 服务器将fileName,target,expire,useLimit,isGroup,file_size, owner记入 file_share_list
// use_cout 默认为0 不需要修改
func ShareFileHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从客户端传入的 JSON 格式参数中获取参数
		var request struct {
			FileName string `json:"fileName"`
			Target   string `json:"target"`
			Expire   int64  `json:"expire"`
			UseLimit int64  `json:"useLimit"`
			IsGroup  bool   `json:"isGroup"`
		}

		if err := c.ShouldBindJSON(&request); err != nil {
			handleError(c, "Parser Error", "Invalid JSON")
			return
		}

		// 从上下文中获取用户信息（例如，从解析的 JWT 中获取用户 ID）
		user, exists := c.Get("user")
		if !exists {
			handleError(c, "Parser Error", "User not found in context")
			return
		}

		// 查询数据库以获取文件的信息（file_size 和 owner）
		var fileOwner string
		var fileSize int64
		err := db.QueryRow("SELECT owner, file_size FROM file_list WHERE file_name = ?", request.FileName).Scan(&fileOwner, &fileSize)
		if err != nil {
			if err == sql.ErrNoRows {
				handleError(c, "DB Error", "File not found")
				return
			}
			handleError(c, "DB Error", "Database query error")
			return
		}
		//确保分享的人存在:
		var target_user string
		err_target := db.QueryRow("select username FROM users where username=?", request.Target).Scan(&target_user)
		if err_target != nil {
			if err == sql.ErrNoRows {
				handleError(c, "DB Error", "target not found")
				return
			}
			handleError(c, "DB Error", "target not found")
			return
		}
		// 确保文件不被分享给文件的拥有者自己
		if user == request.Target {
			handleError(c, "Unauthorized", "Cannot share file with yourself")
			return
		}

		// 查询是否已存在符合条件的记录
		var existingUseCount, existingUseLimit, existingExpire int64
		err = db.QueryRow("SELECT use_count, use_limit, expire FROM file_share_list WHERE file_name = ? AND owner = ? AND target = ?",
			request.FileName, user, request.Target).Scan(&existingUseCount, &existingUseLimit, &existingExpire)
		if err != nil {
			if err == sql.ErrNoRows {
				// 不存在符合条件的记录，执行插入操作
				_, err = db.Exec("INSERT INTO file_share_list (file_name, target, expire, use_limit, is_group, file_size, owner, use_count) VALUES (?, ?, ?, ?, ?, ?, ?, 0)",
					request.FileName, request.Target, request.Expire, request.UseLimit, Utils.ConvertBoolToInt(request.IsGroup), fileSize, fileOwner)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"message": "Error sharing file"})
					return
				}
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "Database query error"})
				return
			}
		} else {
			// 存在符合条件的记录，执行更新操作,主要是将 使用数清零
			_, err = db.Exec("UPDATE file_share_list SET use_count = ?, use_limit = ?, expire = ? WHERE file_name = ? AND owner = ? AND target = ?",
				0, existingUseLimit, existingExpire, request.FileName, user, request.Target)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "Error updating file share"})
				return
			}
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "File shared successfully"})
	}
}
