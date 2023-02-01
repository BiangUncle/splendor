package service

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"splendor/model"
	"splendor/utils"
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
	_ = connectStatus
	table := model.GetDefaultTable()

	curPlayer := table.CurrentPlayer
	c.JSON(http.StatusOK, gin.H{
		"current_player_id":   curPlayer.PlayerID,
		"current_player_name": curPlayer.Name,
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
	_ = connectStatus
	table := model.GetDefaultTable()

	curPlayer := table.NextTurn()

	c.JSON(http.StatusOK, gin.H{
		"current_player_id":   curPlayer.PlayerID,
		"current_player_name": curPlayer.Name,
	})
}

// KeepALive 保活
func KeepALive(c *gin.Context) {
	sessionID, err := GetSessionID(c)
	if err != nil {
		BuildErrorResponse(c, err)
		return
	}

	connectStatus := SessionsMap[sessionID]
	playerID := connectStatus.PlayerID

	model.KeepALive(playerID)

	c.String(http.StatusOK, "ok")
}

// TakeThreeTokens 拿走三个宝石
func TakeThreeTokens(c *gin.Context) {
	sessionID, err := GetSessionID(c)
	if err != nil {
		BuildErrorResponse(c, err)
		return
	}

	connectStatus := SessionsMap[sessionID]
	table := model.GetDefaultTable()
	player, err := model.GetGlobalPlayer(connectStatus.PlayerID)
	if err != nil {
		BuildErrorResponse(c, err)
		return
	}

	tokensStrings := c.Query("tokens")
	var tokens []int
	for i := 0; i < len(tokensStrings); i++ {
		tokens = append(tokens, int(tokensStrings[i]-'0'))
	}

	ret, err := ActionTakeThreeTokens(player, table, tokens)
	if err != nil {
		BuildErrorResponse(c, err)
		return
	}

	nextPlayer := table.NextTurn()

	c.JSON(http.StatusOK, gin.H{
		"ret":              ret,
		"next_player_name": nextPlayer.Name,
	})
}

// TakeDoubleTokens 拿走两个宝石
func TakeDoubleTokens(c *gin.Context) {
	sessionID, err := GetSessionID(c)
	if err != nil {
		BuildErrorResponse(c, err)
		return
	}

	connectStatus := SessionsMap[sessionID]
	table := model.GetDefaultTable()
	player, err := model.GetGlobalPlayer(connectStatus.PlayerID)
	if err != nil {
		BuildErrorResponse(c, err)
		return
	}

	tokenIDString := c.Query("token_id")
	tokenID := utils.ToInt(tokenIDString)
	ret, err := ActionTakeDoubleTokens(player, table, tokenID)
	if err != nil {
		BuildErrorResponse(c, err)
		return
	}

	nextPlayer := table.NextTurn()

	c.JSON(http.StatusOK, gin.H{
		"ret":              ret,
		"next_player_name": nextPlayer.Name,
	})
}

// ReturnTokens 返还多余的宝石
func ReturnTokens(c *gin.Context) {
	sessionID, err := GetSessionID(c)
	if err != nil {
		BuildErrorResponse(c, err)
		return
	}

	connectStatus := SessionsMap[sessionID]
	table := model.GetDefaultTable()
	player, err := model.GetGlobalPlayer(connectStatus.PlayerID)
	if err != nil {
		BuildErrorResponse(c, err)
		return
	}

	tokensStrings := c.Query("tokens")
	var tokenIdx []int
	for i := 0; i < len(tokensStrings); i++ {
		tokenIdx = append(tokenIdx, int(tokensStrings[i]-'0'))
	}

	err = ActionReturnTokens(player, table, tokenIdx)
	if err != nil {
		BuildErrorResponse(c, err)
		return
	}
	c.String(http.StatusOK, "ok")
}
