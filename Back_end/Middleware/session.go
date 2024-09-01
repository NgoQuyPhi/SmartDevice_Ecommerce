package middleware

import (
	database "PJ/SmartDevice_Ecomerce/Back_end/Database"
	models "PJ/SmartDevice_Ecomerce/Back_end/Models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func LoginHandle(c *gin.Context) {
	var LoginData models.LoginData

	err := c.Bind(&LoginData)

	if err != nil {
		c.HTML(201, "notice.html", gin.H{
			"notice": err,
		})
		return
	}

	if LoginData.Username == "" || LoginData.Pass == "" {
		c.HTML(401, "notice.html", gin.H{
			"notice": "username or password is invalid",
		})
		return
	}

	var userdata models.User
	err = database.Instance.
		Table("users").
		Select("*").
		Where("username = ?", LoginData.Username).
		Scan(&userdata).
		Error

	if err != nil {
		c.HTML(201, "notice.html", gin.H{
			"notice": err,
		})
		return
	}

	IsAuthenticated, err := userdata.CheckPassword(LoginData.Pass)

	if err != nil {
		c.HTML(401, "notice.html", gin.H{
			"notice": err,
		})
		return
	}

	if IsAuthenticated {
		session := sessions.Default(c)

		session.Set("ID", userdata.UserId)
		session.Set("name", userdata.Username)
		session.Set("role", userdata.Role)

	}
	c.Redirect(200, "/")
}
