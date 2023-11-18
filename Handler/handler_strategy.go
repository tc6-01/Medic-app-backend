package Handler

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
			handleError(c, "Parser Error", "User not found in context")
			return
		}
		/*
			从数据库中获取用户ID并将策略写进数据库中
		*/
		if err := db.QueryRow("INSERT INTO strategy(strategy_name,strategy_desc,strategy_rule,user_id) VALUES($1,$2,$3,$4)",
			strategyRequest.StrategyName, strategyRequest.StrategyDesc, strategyRequest.StrategyRule, user).Scan(&strategyRequest.StrategyName); err != nil {
			handleError(c, "DB Error", "stragety write error")
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "Stragety Crated successfully"})
		return
	}

}
