package service

import (
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

// BuildErrorResponse 构建错误回复
func BuildErrorResponse(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"msg": err.Error(),
	})
}

// GetSessionID 获取sessionID
func GetSessionID(c *gin.Context) (string, error) {
	session := sessions.Default(c)

	username := session.Get("username")
	var result string

	if username != nil && username != "" {
		result = username.(string)
		if _, ok := SessionsMap[result]; !ok {
			result = "no exist"
			// todo: 需要修改为其他状态码
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

// _addSlash 增加斜杠
func _addSlash(name string) string {
	return "/" + name
}
