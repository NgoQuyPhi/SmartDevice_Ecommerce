package handlers

import (
	database "PJ/SmartDevice_Ecomerce/Back_end/Database"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var SignupData struct {
		Username string `json:"username" form:"username"`
		Password string `json:"password" form:"password"`
		Name     string `json:"name" form:"name"`
		Email    string `json:"email" form:"email"`
		Phone    string `json:"phone" form:"phone"`
	}

	err := c.ShouldBind(&SignupData)

	if err != nil {
		c.HTML(400, "notice.html", gin.H{
			"notice": err,
		})
		return
	}

	err = database.
		Instance.
		Table("users").
		Create(&SignupData).
		Error

	if err != nil {
		c.HTML(400, "notice.html", gin.H{
			"notice": err,
		})
		return
	}

	c.HTML(200, "notice.html", gin.H{
		"notice": "SignUp Success",
	})

}
