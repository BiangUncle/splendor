package service

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"splendor/model"
)

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
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"tableInfo": table.ShowTableInfo(),
	})
	return
}
