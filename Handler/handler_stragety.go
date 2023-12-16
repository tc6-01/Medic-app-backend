package Handler

import (
	"database/sql"
	"net/http"
	"yiliao/Dao"

	"github.com/gin-gonic/gin"
)

/*
创建 & 更新策略
*/
func CreateStragety(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		/*
			构建策略请求对象
				策略名称：string
				策略描述: string
				策略规则: json
		*/
		var strategyRequest struct {
			StrategyName string `json:"strategy_name"`
			StrategyDesc string `json:"strategy_desc"`
			StrategyRule string `json:"strategy_rule"`
		}
		if err := c.BindJSON(&strategyRequest); err != nil {
			handleError(c, "Parser Error", "param  not found")
			return
		}
		// 从上下文中获取用户信息（例如，从解析的 JWT 中获取用户 ID）
		user, exists := c.Get("user")
		if !exists {
			handleError(c, "Parser Error", "User not login!")
			return
		}
		var loginUser *Dao.User
		loginUser, _ = user.(*Dao.User)
		/*
			从数据库中获取用户ID并将策略写进数据库中
		*/
		_, err := db.Exec("INSERT INTO strategy(strategy_name,strategy_desc,strategy_rule,user_id) VALUES($1,$2,$3,$4)",
			strategyRequest.StrategyName, strategyRequest.StrategyDesc, strategyRequest.StrategyRule, loginUser.UserId)
		if err != nil {
			handleError(c, "DB Error", "stragety write error")
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "Stragety Crated successfully"})
		return
	}

}

/*
查看该用户创建的策略
*/
func ListStragety(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从上下文中获取用户信息（例如，从解析的 JWT 中获取用户 ID）
		user, exists := c.Get("user")
		if !exists {
			handleError(c, "Parser Error", "User not login!")
			return
		}
		var loginUser *Dao.User
		loginUser, _ = user.(*Dao.User)
		/*
			从数据库中获取用户ID并将策略写进数据库中
		*/
		rows, err := db.Query("SELECT strategy_name,strategy_desc,strategy_rule,FROM strategy WHERE user_id = ?", loginUser.UserId)

		if err != nil {
			handleError(c, "DB Error", "Database query error")
			return
		}
		defer rows.Close()
		// 创建一个切片来存储用户创建的策略
		var strageties []Dao.Stragety
		// 遍历查询结果并填充 sharedFiles 切片
		for rows.Next() {
			var (
				name   string
				desc   string
				expire int64
			)

			err := rows.Scan(&name, &desc, &expire)
			if err != nil {
				handleError(c, "DB Error", "Error scanning rows")
				return
			}

			// 创建共享文件对象并填充数据
			stragety := Dao.Stragety{
				Name:   name,
				Desc:   desc,
				Expire: expire,
			}

			strageties = append(strageties, stragety)
		}

		if err := rows.Err(); err != nil {
			handleError(c, "DB Error", "Error in rows")
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "Stragety Crated successfully", "data": strageties})
		return
	}
}

/*
删除策略
*/
func RemoveStragety(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		/*
			构建策略请求对象
				策略名称：string
				策略描述: string
				策略规则: json
		*/
		var strategyRequest struct {
			StrategyId int64 `json:"strategy_id"`
		}
		if err := c.BindJSON(&strategyRequest); err != nil {
			handleError(c, "Parser Error", "param  not found")
			return
		}
		// 从上下文中获取用户信息（例如，从解析的 JWT 中获取用户 ID）
		user, exists := c.Get("user")
		if !exists {
			handleError(c, "Parser Error", "User not login!")
			return
		}
		var loginUser *Dao.User
		loginUser, _ = user.(*Dao.User)
		/*
			根据用户ID与策略ID删除策略
		*/
		_, err := db.Exec("UPDATE  strategy set is_delete = 1 where user_id = $1 and strategy_id = $2  VALUES($1,$2,$3,$4)",
			loginUser.UserId, strategyRequest.StrategyId)
		if err != nil {
			handleError(c, "DB Error", "stragety write error")
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "Stragety remove successfully"})
		return
	}
}
