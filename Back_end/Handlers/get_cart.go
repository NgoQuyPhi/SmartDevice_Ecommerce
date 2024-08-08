package handlers

import (
	models "PJ/SmartDevice_Ecomerce/Back_end/Models"

	"github.com/gin-gonic/gin"
)

func GetCart(c *gin.Context) {
	cart := models.GetCartFromContext(c)
	c.HTML(200, "cart.html", gin.H{
		"cart": cart,
	})
}
