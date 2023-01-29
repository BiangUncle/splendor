package service

import (
	"errors"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
	"splendor/model"
	"splendor/utils"
	"time"
)

type ConnectStatus struct {
	CreateTime time.Time // 连接创建时间
	TableID    string    // 桌台ID
	PlayerID   string    // 玩家ID
}

func (c *ConnectStatus) Info() string {
	info := ""
	info = info + fmt.Sprintf("CreateTime: %+v\n", c.CreateTime)
	info = info + fmt.Sprintf("TableID: %+v\n", c.TableID)
	return info
}

var SessionsMap = make(map[string]*ConnectStatus)

/*
有空看看 https://blog.csdn.net/qq_43716830/article/details/124431938

游戏服务端框架 leaf
https://github.com/name5566/leaf/blob/master/TUTORIAL_ZH.md
*/

func BuildErrorResponse(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"msg": err.Error(),
	})
}

func Ping(c *gin.Context) {
	c.String(http.StatusOK, "Pong")
}

func Join(c *gin.Context) {

	// 获取用户名
	username := c.Query("username")
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
		"username":   username,
		"msg":        "set session success",
		"player_num": 1,
		"table_id":   tableID,
		"player_id":  playerID,
		"session_id": uuid,
	})
}

func Leave(c *gin.Context) {

	session := sessions.Default(c)

	uuid := session.Get("username")
	if uuid == nil || uuid.(string) == "" {
		c.String(http.StatusOK, "user has existed")
		return
	}

	session.Delete("username")

	delete(SessionsMap, uuid.(string))

	c.String(http.StatusOK, "delete session success")
}

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

func GetSessionID(c *gin.Context) (string, error) {
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
			return result, errors.New("no sessionID")
		}
	} else {
		result = "no exist"
		c.JSON(http.StatusOK, gin.H{
			"result": result,
		})
		return result, errors.New("no sessionID")
	}

	return result, nil
}

// Run 运行服务端
func Run() {
	e := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	e.Use(sessions.Sessions("SessionID", store))
	e.GET("/ping", Ping)
	e.GET("/join", Join)
	e.GET("/leave", Leave)
	e.GET("/alive", Alive)
	e.GET("/table_info", TableInfo)
	e.GET("/cur_player", AskWhichTurn)
	e.GET("/next_turn", NextTurn)

	e.Run(":8765")
}
