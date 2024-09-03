package handlers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	session := sessions.Default(c)

	session.Clear()
	session.Set("isauthenticated", false)
	session.Save()
	c.HTML(200, "notice.html", gin.H{
		"notice": "Logout Success",
	})

}
