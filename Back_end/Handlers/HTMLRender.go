package handlers

import "github.com/gin-gonic/gin"

func ShowLoginPage(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}
func ShowSignupPage(c *gin.Context) {
	c.HTML(200, "register.html", nil)
}
