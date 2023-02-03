package service

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"splendor/model"
	"splendor/utils"
	"time"
)

// Ping Ping
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "Pong")
}

// Join 加入游戏
func Join(c *gin.Context) {

	// 获取用户名
	username := c.Query("username")
	_ = username
	session := sessions.Default(c)

	uuid := utils.GetUuidV1()

	session.Set("username", uuid)

	err := session.Save()
	if err != nil {
		BuildErrorResponse(c, err)
		return
	}

	player, playerID, err := model.JoinNewPlayer()
	if err != nil {
		BuildErrorResponse(c, err)
		return
	}

	_, tableID, err := model.JoinDefaultTable(player)
	if err != nil {
		BuildErrorResponse(c, err)
		return
	}

	SessionsMap[uuid] = &ConnectStatus{
		CreateTime: time.Now(),
		TableID:    tableID,
		PlayerID:   playerID,
	}

	fmt.Println(SessionsMap[uuid].Info())

	c.JSON(http.StatusOK, gin.H{
		"status":     "created",
		"username":   player.Name,
		"msg":        "set session success",
		"player_num": 1,
		"table_id":   tableID,
		"player_id":  playerID,
		"session_id": uuid,
	})
}

// Leave 离开游戏
func Leave(c *gin.Context) {

	sessionID, err := GetSessionID(c)
	if err != nil {
		BuildErrorResponse(c, err)
		return
	}

	session := SessionsMap[sessionID]
	model.LeaveDefaultTable(session.PlayerID)  // 清理桌面中的角色
	model.DeleteGlobalPlayer(session.PlayerID) // 清理全局角色对象

	delete(SessionsMap, sessionID)

	c.String(http.StatusOK, "delete session success")
}

// Alive 心跳
func Alive(c *gin.Context) {
	session := sessions.Default(c)

	username := session.Get("username")
	var result string

	if username != nil && username != "" {
		result = username.(string)
		if _, ok := SessionsMap[result]; !ok {
			result = "no exist"
		}
	} else {
		result = "no exist"
	}

	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

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

	ret, err := model.ActionTakeThreeTokens(player, table, tokens)
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
	ret, err := model.ActionTakeDoubleTokens(player, table, tokenID)
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

	err = model.ActionReturnTokens(player, table, tokenIdx)
	if err != nil {
		BuildErrorResponse(c, err)
		return
	}
	c.String(http.StatusOK, "ok")
}

// GetCurrentSessionTableAndPlayer 获取当前session所在的桌台和玩家
func GetCurrentSessionTableAndPlayer(c *gin.Context) (*model.Table, *model.Player, error) {
	sessionID, err := GetSessionID(c)
	if err != nil {
		BuildErrorResponse(c, err)
		return nil, nil, err
	}

	connectStatus := SessionsMap[sessionID]
	table := model.GetDefaultTable()
	player, err := model.GetGlobalPlayer(connectStatus.PlayerID)
	if err != nil {
		BuildErrorResponse(c, err)
		return nil, nil, err
	}
	return table, player, nil
}

// PurchaseDevelopmentCardByTokens 使用 tokens 购买发展卡
func PurchaseDevelopmentCardByTokens(c *gin.Context) {
	table, player, err := GetCurrentSessionTableAndPlayer(c)
	if err != nil {
		BuildErrorResponse(c, err)
		return
	}

	tokensStrings := c.Query("tokens")
	cardIdx := model.DevelopmentCardIndexTransfer(utils.ToInt(c.Query("card_idx")))
	var tokens []int
	for i := 0; i < len(tokensStrings); i++ {
		tokens = append(tokens, int(tokensStrings[i]-'0'))
	}

	tokens, err = model.IntList2TokenStack(tokens)
	if err != nil {
		BuildErrorResponse(c, err)
		return
	}

	err = model.PurchaseDevelopmentCard(player, tokens, table, cardIdx)
	if err != nil {
		BuildErrorResponse(c, err)
		return
	}

	err = model.ReceiveNoble(player, table)
	if err != nil {
		BuildErrorResponse(c, err)
		return
	}
	nextPlayer := table.NextTurn()

	c.JSON(http.StatusOK, gin.H{
		"next_player_name": nextPlayer.Name,
	})
}

// ReserveDevelopmentCard 保存发展卡
func ReserveDevelopmentCard(c *gin.Context) {
	table, player, err := GetCurrentSessionTableAndPlayer(c)
	if err != nil {
		BuildErrorResponse(c, err)
		return
	}

	cardIdx := model.DevelopmentCardIndexTransfer(utils.ToInt(c.Query("card_idx")))

	err = model.ReserveDevelopmentCard(player, table, cardIdx)
	if err != nil {
		BuildErrorResponse(c, err)
		return
	}
	err = model.ReceiveNoble(player, table)
	if err != nil {
		BuildErrorResponse(c, err)
		return
	}

	nextPlayer := table.NextTurn()

	c.JSON(http.StatusOK, gin.H{
		"next_player_name": nextPlayer.Name,
	})
}

// PurchaseHandCard 购买手卡
func PurchaseHandCard(c *gin.Context) {
	table, player, err := GetCurrentSessionTableAndPlayer(c)
	if err != nil {
		BuildErrorResponse(c, err)
		return
	}

	tokensStrings := c.Query("tokens")
	cardIdx := model.DevelopmentCardIndexTransfer(utils.ToInt(c.Query("card_idx")))

	var tokens []int
	for i := 0; i < len(tokensStrings); i++ {
		tokens = append(tokens, int(tokensStrings[i]-'0'))
	}

	tokens, err = model.IntList2TokenStack(tokens)
	if err != nil {
		BuildErrorResponse(c, err)
		return
	}

	err = model.PurchaseHandCard(player, tokens, table, cardIdx)
	if err != nil {
		BuildErrorResponse(c, err)
		return
	}

	err = model.ReceiveNoble(player, table)
	if err != nil {
		BuildErrorResponse(c, err)
		return
	}
	nextPlayer := table.NextTurn()

	c.JSON(http.StatusOK, gin.H{
		"next_player_name": nextPlayer.Name,
	})
}
