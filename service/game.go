package service

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"splendor/model"
)

// TableInfo 获取桌面信息HTTP接口
func TableInfo(c *gin.Context) {
	session := sessions.Default(c)

	username := session.Get("username")
	var result string

	if username != nil && username != "" {
		result = username.(string)
		if _, ok := SessionsMap[result]; !ok {
			result = "no exist"
			c.JSON(http.StatusOK, gin.H{
				"result": result,
			})
			return
		}
	} else {
		result = "no exist"
		c.JSON(http.StatusOK, gin.H{
			"result": result,
		})
		return
	}

	connectStatus := SessionsMap[result]
	table, err := model.GetGlobalTable(connectStatus.TableID)
	if err != nil {
		BuildErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tableInfo": table.ShowTableInfo(),
	})
	return
}

// AskWhichTurn 查询当前玩家轮次接口
func AskWhichTurn(c *gin.Context) {
	sessionID, err := GetSessionID(c)
	if err != nil {
		BuildErrorResponse(c, err)
		return
	}

	connectStatus := SessionsMap[sessionID]
	table, err := model.GetGlobalTable(connectStatus.TableID)
	if err != nil {
		BuildErrorResponse(c, err)
		return
	}

	curPlayer := table.CurrentPlayer
	c.JSON(http.StatusOK, gin.H{
		"current_player_id": curPlayer.PlayerID,
	})
}

// NextTurn 执行下一个轮次接口
func NextTurn(c *gin.Context) {
	sessionID, err := GetSessionID(c)
	if err != nil {
		BuildErrorResponse(c, err)
		return
	}

	connectStatus := SessionsMap[sessionID]
	table, err := model.GetGlobalTable(connectStatus.TableID)
	if err != nil {
		BuildErrorResponse(c, err)
		return
	}

	curPlayer := table.NextTurn()

	fmt.Println(curPlayer.PlayerInfoString())
	c.JSON(http.StatusOK, gin.H{
		"current_player_id":   curPlayer.PlayerID,
		"current_player_name": curPlayer.Name,
	})
}
