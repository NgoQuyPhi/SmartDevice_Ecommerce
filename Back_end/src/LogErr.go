package src

import "github.com/gin-gonic/gin"

func LogErr(c *gin.Context, errCode int, notice string) {
	c.HTML(errCode, "notice.html", gin.H{
		"notice": notice,
	})
}
