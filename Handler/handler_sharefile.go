package Handler

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"yiliao/Dao"
)

// ShareFileHandler 指定共享策略进行病历共享
func ShareFileHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("user")
		if !exists {
			handleError(c, "Unauthorized", "当前用户未登录")
			return
		}
		d := user.(*Dao.User)
		// 从客户端传入的 JSON 格式参数中获取参数
		var request struct {
			FileName string `json:"fileName"`
			Target   string `json:"target"`
			Expire   int64  `json:"expire"`
			UseLimit int64  `json:"useLimit"`
			Name     string `json:"name"`
			Desc     string `json:"desc"`
		}

		if err := c.ShouldBindJSON(&request); err != nil {
			handleError(c, "Parser Error", "输入参数有误")
			return
		}
		// 查询数据库以获取文件的信息（file_owner 与 file_id）
		var ownerId, fileId int64
		err := db.QueryRow("SELECT owner_id,file_id FROM files WHERE file_name = ?", request.FileName).Scan(&ownerId, &fileId)
		if err != nil {
			handleError(c, "DB Error", "查询失败")
			return
		}
		//确保分享的人存在
		var targetId int64
		errTarget := db.QueryRow("select user_id FROM user where user_name=?", request.Target).Scan(&targetId)
		if errTarget != nil {
			handleError(c, "DB Error", "目标用户不存在")
			return
		}
		// 查询是否已存在符合条件的记录
		var sId int64
		err = db.QueryRow(`SELECT id FROM share_files 
          WHERE fileId = ? AND target_user_id = ? `,
			fileId, targetId).Scan(&sId)
		if err != nil {
			if errors.Is(sql.ErrNoRows, err) {
				// 不存在符合条件的记录，执行插入操作
				_, err = db.Exec(`INSERT INTO share_files 
    				(fileId, name, des, expire, from_user_id, target_user_id, use_limit) VALUES (?,?,?,?,?,?,?)`,
					fileId, request.Name, request.Desc, request.Expire, d.UserId, targetId, request.UseLimit)
				if err != nil {
					handleError(c, "DB Error", "插入失败")
					return
				}
			} else {
				handleError(c, "DB Error", "查询失败")
				return
			}
		} else {
			// 存在符合条件的记录，执行更新操作,覆盖原来的策略,两个用户之间同一病历只能有一种共享策略
			_, err = db.Exec(`UPDATE share_files 
			SET name=?, desc=?, expire=?, use_limit=?,use_count = ? WHERE id = ?`,
				request.Name, request.Desc, request.Expire, request.UseLimit, 0, sId)
			if err != nil {
				handleError(c, "DB Error", "更新失败")
				return
			}
		}
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "病历已共享"})
	}
}
