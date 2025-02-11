package src

import (
	models "PJ/SmartDevice_Ecomerce/Back_end/Models"

	"github.com/gin-gonic/gin"
)

func GetCart(c *gin.Context) {
	cart := models.GetCartFromContext(c)
	c.HTML(200, "viewCart.html", gin.H{
		"cart": cart,
	})
}
