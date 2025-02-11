package src

import (
	models "PJ/SmartDevice_Ecomerce/Back_end/Models"
	repositories "PJ/SmartDevice_Ecomerce/Back_end/Repositories"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var data models.SignupData

	err := c.ShouldBind(&data)

	if err != nil {
		c.HTML(400, "notice.html", gin.H{
			"notice": err,
		})
		return
	}

	err = repositories.
		Instance.
		Table("users").
		Create(&data).
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
