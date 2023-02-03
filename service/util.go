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
		}
	} else {
		result = "no exist"
	}

	if result != "no exist" {
		return result, nil
	}

	result = c.Request.Header.Get("cookie")
	if result != "" {
		return result, nil
	}

	result = c.Query("session_id")
	if result != "" {
		return result, nil
	}

	result = "no exist"
	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
	return result, errors.New("no sessionID")

}

// _addSlash 增加斜杠
func _addSlash(name string) string {
	return "/" + name
}
