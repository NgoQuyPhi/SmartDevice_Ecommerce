package handlers

import (
	models "PJ/SmartDevice_Ecomerce/Back_end/Models"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Add_item_quantity(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}
	session := sessions.Default(c)
	sessionID := session.Get("ID")
	cart := models.GetCartFromContext(c)

	cart.AddItem(productID, 1)
	models.UpdateCart(sessionID, cart)

	c.Redirect(http.StatusFound, "/cart")
}
func Subtract_item_quantity(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}
	session := sessions.Default(c)
	sessionID := session.Get("ID")
	cart := models.GetCartFromContext(c)

	cart.AddItem(productID, -1)
	models.UpdateCart(sessionID, cart)

	c.Redirect(http.StatusFound, "/cart")
}
