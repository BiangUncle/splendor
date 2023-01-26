package service

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
	"splendor/utils"
	"time"
)

type ConnectStatus struct {
	CreateTime time.Time
}

func (c *ConnectStatus) Info() string {
	info := ""
	info = info + fmt.Sprintf("CreateTime: %+v\n", c.CreateTime)
	return info
}

var SessionsMap = make(map[string]*ConnectStatus)

/*
有空看看 https://blog.csdn.net/qq_43716830/article/details/124431938

游戏服务端框架 leaf
https://github.com/name5566/leaf/blob/master/TUTORIAL_ZH.md
*/

// Count 统计在线人数
func Count(c *gin.Context) {
	session := sessions.Default(c)
	var count int
	v := session.Get("count")
	if v == nil {
		count = 0
	} else {
		count = v.(int)
		count++
	}
	session.Set("count", count)
	session.Save()
	c.JSON(200, gin.H{"count": count})
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
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "set session failed",
		})
		return
	}

	SessionsMap[uuid] = &ConnectStatus{
		CreateTime: time.Now(),
	}
	fmt.Println(SessionsMap[uuid].Info())

	c.JSON(http.StatusOK, gin.H{
		"status":     "created",
		"username":   username,
		"msg":        "set session success",
		"player_num": 1,
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

// Run 运行服务端
func Run() {
	e := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	e.Use(sessions.Sessions("SessionID", store))
	e.GET("/ping", Ping)
	e.GET("/join", Join)
	e.GET("/leave", Leave)
	e.GET("/alive", Alive)
	e.Run(":8765")
}
