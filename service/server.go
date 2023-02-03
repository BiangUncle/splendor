package service

/*
有空看看 https://blog.csdn.net/qq_43716830/article/details/124431938

游戏服务端框架 leaf
https://github.com/name5566/leaf/blob/master/TUTORIAL_ZH.md
*/

import (
	"context"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
	"splendor/consts"
	"splendor/model"
	"splendor/utils"
	"time"
)

var SessionsMap = make(map[string]*ConnectStatus)

// ConnectStatus 连接状态
type ConnectStatus struct {
	CreateTime time.Time // 连接创建时间
	TableID    string    // 桌台ID
	PlayerID   string    // 玩家ID
}

// Info 连接状态展示
func (c *ConnectStatus) Info() string {
	info := ""
	info = info + fmt.Sprintf("CreateTime: %+v\n", c.CreateTime)
	info = info + fmt.Sprintf("TableID: %+v\n", c.TableID)
	return info
}

// InitHandler 初始化handler
func InitHandler(e *gin.Engine) *gin.Engine {
	e.GET(_addSlash(consts.Ping), Ping)
	e.GET(_addSlash(consts.Join), Join)
	e.GET(_addSlash(consts.Leave), Leave)
	e.GET(_addSlash(consts.Alive), Alive)
	e.GET(_addSlash(consts.TableInfo), TableInfo)
	e.GET(_addSlash(consts.AskWhichTurn), AskWhichTurn)
	e.GET(_addSlash(consts.NextTurn), NextTurn)
	e.GET(_addSlash(consts.KeepALive), KeepALive)
	e.GET(_addSlash(consts.TakeThreeTokens), TakeThreeTokens)
	e.GET(_addSlash(consts.TakeDoubleTokens), TakeDoubleTokens)
	e.GET(_addSlash(consts.ReturnTokens), ReturnTokens)
	e.GET(_addSlash(consts.PurchaseDevelopmentCardByTokens), PurchaseDevelopmentCardByTokens)
	e.GET(_addSlash(consts.ReserveDevelopmentCardApi), ReserveDevelopmentCard)
	e.GET(_addSlash(consts.PurchaseHandCardApi), PurchaseHandCard)
	return e
}

// RunServer 运行服务器
func RunServer(c *OneCron, srv *http.Server) {
	c.wg.Add(1)
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		fmt.Println(err)
	}
	fmt.Println("服务端: ", "程序退出")
	c.wg.Done()
}

// RunCron 运行守护线程
func RunCron(c *OneCron) {
	c.wg.Add(1)
	c.Run()
	c.wg.Done()
}

// ShutDown 关闭服务
func ShutDown(c *OneCron, srv *http.Server) {
	err := srv.Shutdown(context.Background()) // 关闭服务器
	if err != nil {
		fmt.Println(err)
	}
	c.Stop() // 关闭清理线程
	fmt.Println("服务端: ", "等待其他协程退出")
	c.wg.Wait()
}

// Input 服务器接受输入
func Input() {
	for {
		s, _ := utils.InputString("")
		if s == "table" {
			fmt.Println(model.GetDefaultTable().ShowTableInfo())
			continue
		}
		break
	}
}

func SetSession(e *gin.Engine) *gin.Engine {
	store := cookie.NewStore([]byte("secret"))
	e.Use(sessions.Sessions("SessionID", store))
	return e
}

func NewServer(addr string, h http.Handler) *http.Server {
	return &http.Server{
		Addr:    ":8765",
		Handler: h,
	}
}

// Run 启动!
func Run() {
	e := gin.Default()
	e = SetSession(e)
	e = InitHandler(e)

	srv := NewServer(":8765", e)
	c := ConstructCron()

	go RunServer(c, srv)
	go RunCron(c)

	stop := make(chan struct{})
	defer close(stop)

	Input()

	ShutDown(c, srv)

}
