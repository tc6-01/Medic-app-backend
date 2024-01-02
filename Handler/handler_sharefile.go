package Handler

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"yiliao/Dao"

	"github.com/gin-gonic/gin"
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
			IsAllow  int64  `json:"isAllow"`
			State    string `json:"state"`
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
		var dId, useCount, useLimit int64
		flag := true
		// 由当前用户共享的，找到共享目标用户为当前用户的共享记录进行更新，如果没有就不更新（因为如果该病历为自己的，就不需要更新）
		if request.State == "fromShared" || request.State == "shared" {
			query := `select id, use_count,use_limit from share_files where target_user_id = ? and fileId = ?`
			err := db.QueryRow(query, d.UserId, fileId).Scan(&dId, &useCount, &useLimit)
			if err != nil {
				if errors.Is(sql.ErrNoRows, err) {
					log.Println("当前为病历所有者进行文件共享，不需要更新可访问次数")
					flag = false
				} else {
					handleError(c, "DB Error", "数据库查询错误")
					return
				}
			}
			if flag {
				if (useLimit - useCount) < request.UseLimit {
					handleError(c, "Unauthorized", "当前访问次数超出原有权限")
					return
				}
				// 执行原有记录可访问次数更新
				_, err = db.Exec(`UPDATE share_files SET use_count = ? WHERE id = ?`, useCount+request.UseLimit, dId)
				if err != nil {
					handleError(c, "DB Error", "可访问次数更新失败")
					return
				}
			}
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
    				(fileId, name, des, expire, from_user_id, target_user_id, use_limit,is_allow) VALUES (?,?,?,?,?,?,?,?)`,
					fileId, request.Name, request.Desc, request.Expire, d.UserId, targetId, request.UseLimit, request.IsAllow)
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
			SET name=?, des=?, expire=?, use_limit=?,use_count = ? ,is_allow = ? WHERE id = ?`,
				request.Name, request.Desc, request.Expire, request.UseLimit, 0, request.IsAllow, sId)
			if err != nil {
				handleError(c, "DB Error", "更新失败")
				return
			}
		}
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "病历已共享"})
	}
}
