package src

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GetSessionData(c *gin.Context) (interface{}, interface{}, interface{}) {
	session := sessions.Default(c)
	IsAuthenticated := session.Get("isauthenticated")

	name := session.Get("name")
	role := session.Get("role")

	return name, role, IsAuthenticated

}
